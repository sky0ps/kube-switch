package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Constants for context types (based on name patterns)
const (
	typeDevelopment = "development"
	typeStaging     = "staging"
	typeProduction  = "production"
	typeUnknown     = "unknown"
)

// getKubeConfig loads the Kubernetes configuration
func getKubeConfig() (*api.Config, string, error) {
	// Get kubeconfig path, using KUBECONFIG env var or default location
	kubeConfigPath := os.Getenv("KUBECONFIG")
	if kubeConfigPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, "", fmt.Errorf("failed to get home directory: %w", err)
		}
		kubeConfigPath = filepath.Join(homeDir, ".kube", "config")
	}

	// Load kubeconfig
	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	return config, kubeConfigPath, nil
}

// saveKubeConfig writes the Kubernetes configuration back to disk
func saveKubeConfig(config *api.Config, path string) error {
	return clientcmd.WriteToFile(*config, path)
}

// determineContextType tries to determine if a context is dev, staging, or prod
func determineContextType(name string) string {
	nameLower := strings.ToLower(name)

	if strings.Contains(nameLower, "prod") || strings.Contains(nameLower, "prd") {
		return typeProduction
	} else if strings.Contains(nameLower, "stage") || strings.Contains(nameLower, "stg") {
		return typeStaging
	} else if strings.Contains(nameLower, "dev") || strings.Contains(nameLower, "development") {
		return typeDevelopment
	}

	return typeUnknown
}

// getContextColor returns the color for a context based on its type
func getContextColor(contextType string) tcell.Color {
	switch contextType {
	case typeProduction:
		return tcell.ColorDarkMagenta // Deep purple for production
	case typeStaging:
		return tcell.ColorPurple // Purple for staging
	case typeDevelopment:
		return tcell.ColorBlue // Blue for development
	default:
		return tcell.ColorLightBlue // Light blue for unknown
	}
}

// getContextInfo returns formatted context info
func getContextInfo(context *api.Context) string {
	if context == nil {
		return "No context information available"
	}

	return fmt.Sprintf("Cluster: %s | User: %s | Namespace: %s",
		context.Cluster,
		context.AuthInfo,
		context.Namespace)
}

// getNamespaces returns the list of namespaces in the context
// Note: In a real implementation, this would query the k8s API
// For this demo, we'll just return some dummy namespaces
func getNamespaces(contextName string) []string {
	// In a real implementation, we would use the Kubernetes API to get actual namespaces
	// For demo purposes, we're returning dummy namespaces
	return []string{
		"default",
		"kube-system",
		"kube-public",
		"kube-node-lease",
		"monitoring",
		"logging",
		"app-frontend",
		"app-backend",
		"database",
	}
}

func main() {
	// Load Kubernetes config
	config, configPath, err := getKubeConfig()
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	// Get current context
	currentContext := config.CurrentContext

	// Create the application
	app := tview.NewApplication()

	// Create the pages to switch between context and namespace views
	pages := tview.NewPages()

	// Create a list for the contexts
	contextList := tview.NewList().
		SetHighlightFullLine(true).
		SetWrapAround(true)

	// Text view for context details
	contextDetails := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Select a context to view details")

	// Add contexts to the list
	var contexts []string
	var contextTypes []string

	for name := range config.Contexts { // Iterate only over names first
		contexts = append(contexts, name)
		contextType := determineContextType(name)
		contextTypes = append(contextTypes, contextType)
	}

	// Add contexts to the list with proper colors
	contextList.Clear()
	for i, name := range contexts {
		displayText := name
		if name == currentContext {
			displayText = "► " + displayText // Mark current context
		}
		ctx := config.Contexts[name] // Get context object
		contextType := contextTypes[i]
		// *** FIX APPLIED HERE: Use the 'ctx' variable ***
		contextList.AddItem(displayText, getContextInfo(ctx), 0, nil).
			SetSelectedTextColor(getContextColor(contextType))
	}

	// Function to switch context
	switchContext := func(contextName string) error {
		// Update config with the new context
		config.CurrentContext = contextName

		// Save the updated config
		if err := saveKubeConfig(config, configPath); err != nil {
			app.Stop()
			log.Fatalf("Error saving kubeconfig: %v", err)
			return err
		}

		// Show a confirmation message
		modal := tview.NewModal().
			SetText(fmt.Sprintf("Switched to context: %s", contextName)).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.Stop() // Exit after switching context
			})

		pages.AddPage("contextConfirm", modal, true, true)
		return nil
	}

	// Update context details when selection changes
	contextList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(contexts) {
			contextName := contexts[index]
			ctx := config.Contexts[contextName] // This ctx is used correctly below
			contextType := contextTypes[index]

			details := fmt.Sprintf("[%s]%s[white]\n\n%s",
				getContextColor(contextType).Name(),
				contextName,
				getContextInfo(ctx)) // Used here

			contextDetails.SetText(details)
		}
	})

	// Function to show namespace selection for a context
	showNamespaceSelection := func(contextName string) {
		// Create a list for namespaces
		namespaceList := tview.NewList().
			SetHighlightFullLine(true).
			SetWrapAround(true)

		// Get current namespace for the context
		currentNamespace := ""
		if ctx, exists := config.Contexts[contextName]; exists && ctx != nil {
			currentNamespace = ctx.Namespace
		}

		// Add namespaces to the list
		namespaces := getNamespaces(contextName)
		for _, ns := range namespaces {
			displayText := ns
			if ns == currentNamespace {
				displayText = "► " + displayText
			}
			namespaceList.AddItem(displayText, "Select to switch to this namespace", 0, nil)
		}

		// Set the selected namespace handler
		namespaceList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
			selectedNamespace := namespaces[index]

			// Update config with the new namespace
			if ctx, exists := config.Contexts[contextName]; exists && ctx != nil {
				ctx.Namespace = selectedNamespace

				// Save the updated config
				if err := saveKubeConfig(config, configPath); err != nil {
					app.Stop()
					log.Fatalf("Error saving kubeconfig: %v", err)
				}

				// Show a confirmation message
				modal := tview.NewModal().
					SetText(fmt.Sprintf("Switched namespace to: %s in context: %s", selectedNamespace, contextName)).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						// Instead of switching back, maybe just remove the modal and keep the context list focused?
						// Or perhaps exit if the goal is just a quick switch? For now, let's exit like context switching.
						app.Stop() // Exit after switching namespace
						// Alternatively, to go back to the context list:
						// pages.RemovePage("namespaceConfirm")
						// pages.SwitchToPage("contexts")
						// app.SetFocus(contextList)
					})

				pages.AddPage("namespaceConfirm", modal, true, true)
			}
		})

		// Create a frame for the namespace list
		nsFrame := tview.NewFrame(namespaceList).
			AddText(fmt.Sprintf("Namespaces for context: %s", contextName), true, tview.AlignCenter, tcell.ColorPurple).
			AddText("↑/↓: Navigate | Enter: Select | Esc: Back | Ctrl-C: Quit", false, tview.AlignCenter, tcell.ColorBlue)

		// Add input capture to handle Escape key
		nsFrame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.RemovePage("namespaces") // Remove the namespace page
				pages.SwitchToPage("contexts") // Switch back to context view
				app.SetFocus(contextList)      // Ensure context list has focus
				return nil
			}
			return event
		})

		// Add the namespace page
		pages.AddPage("namespaces", nsFrame, true, true)
		app.SetFocus(namespaceList) // Focus the namespace list when it appears
	}

	// Set the selected context handler
	contextList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		selectedContext := contexts[index]

		// If switching to production, show confirmation
		if contextTypes[index] == typeProduction {
			modal := tview.NewModal().
				SetText(fmt.Sprintf("Warning: You are switching to a PRODUCTION context: %s\nAre you sure?", selectedContext)).
				AddButtons([]string{"Cancel", "Switch Context", "Switch Namespace"}). // Changed 3rd option label
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					pages.RemovePage("prodConfirm") // Remove the modal
					if buttonIndex == 1 {           // Switch Context
						switchContext(selectedContext) // This will eventually call app.Stop()
					} else if buttonIndex == 2 { // Switch Namespace
						// First switch context if not already current, then show namespaces
						if config.CurrentContext != selectedContext {
							// Update config in memory first
							config.CurrentContext = selectedContext
							// Save the updated config (without exiting yet)
							if err := saveKubeConfig(config, configPath); err != nil {
								app.Stop()
								log.Fatalf("Error saving kubeconfig: %v", err)
								return // Stop further processing on error
							}
							// Update the context list display marker (optional but nice)
							// Need to redraw the list or manually update item text
						}
						showNamespaceSelection(selectedContext)
					} else { // Cancel
						pages.SwitchToPage("contexts")
						app.SetFocus(contextList)
					}
				})

			pages.AddPage("prodConfirm", modal, true, true)
			app.SetFocus(modal) // Focus the modal
		} else {
			// Not production, show namespace options directly
			modal := tview.NewModal().
				SetText(fmt.Sprintf("Context: %s\nWhat would you like to do?", selectedContext)).
				AddButtons([]string{"Cancel", "Switch Context", "Switch Namespace"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					pages.RemovePage("contextOptions") // Remove the modal
					if buttonIndex == 1 {              // Switch Context
						switchContext(selectedContext) // This will eventually call app.Stop()
					} else if buttonIndex == 2 { // Switch Namespace
						// First switch context if not already current, then show namespaces
						if config.CurrentContext != selectedContext {
							// Update config in memory first
							config.CurrentContext = selectedContext
							// Save the updated config (without exiting yet)
							if err := saveKubeConfig(config, configPath); err != nil {
								app.Stop()
								log.Fatalf("Error saving kubeconfig: %v", err)
								return // Stop further processing on error
							}
							// Update the context list display marker (optional but nice)
						}
						showNamespaceSelection(selectedContext)
					} else { // Cancel
						pages.SwitchToPage("contexts")
						app.SetFocus(contextList)
					}
				})

			pages.AddPage("contextOptions", modal, true, true)
			app.SetFocus(modal) // Focus the modal
		}
	})

	// Create layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contextList, 0, 1, true). // Give context list focus initially
		AddItem(contextDetails, 10, 1, false)

	// Create a frame for the context view
	contextFrame := tview.NewFrame(flex).
		AddText("kube switch: Kubernetes Context Switcher", true, tview.AlignCenter, tcell.ColorPurple).
		AddText("↑/↓: Navigate | Enter: Select | Ctrl-C: Quit", false, tview.AlignCenter, tcell.ColorBlue)

	// Add context page
	pages.AddPage("contexts", contextFrame, true, true)

	// Set input handler for Ctrl-C to quit
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Global Ctrl-C handler
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
			return nil
		}
		// Allow other keys to pass through to focused element
		return event
	})

	// Run the application
	if err := app.SetRoot(pages, true).SetFocus(contextList).Run(); err != nil { // Set initial focus
		log.Fatalf("Error running application: %v", err)
	}
}

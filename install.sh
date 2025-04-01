#!/bin/bash

set -e

# Define installation locations
BIN_DIR="$HOME/bin"
INSTALL_PATH="$BIN_DIR/kcs"

# Function for purple-green gradient ASCII art
print_banner() {
    # Define gradient colors (purple to green)
    C1='\033[38;5;93m'  # Light purple
    C2='\033[38;5;92m'  # Purple
    C3='\033[38;5;91m'  # Dark purple
    C4='\033[38;5;90m'  # Purple-magenta
    C5='\033[38;5;54m'  # Dark magenta
    C6='\033[38;5;55m'  # Magenta-blue
    C7='\033[38;5;56m'  # Blue-green
    C8='\033[38;5;49m'  # Light cyan
    C9='\033[38;5;48m'  # Green
    C10='\033[38;5;47m' # Bright green
    NC='\033[0m'       # No Color

    echo -e "${C1}╔╗ ╔═╗  ╦╔═╦ ╦╔╗ ╔═╗  ╔═╗╦ ╦╦╔╦╗╔═╗╦ ╦${NC}"
    echo -e "${C2}╠╩╗║ ║  ╠╩╗║ ║╠╩╗║╣   ╚═╗║║║║ ║ ║  ╠═╣${NC}"
    echo -e "${C3}╚═╝╚═╝  ╩ ╩╚═╝╚═╝╚═╝  ╚═╝╚╩╝╩ ╩ ╚═╝╩ ╩${NC}"
    echo -e "${C4}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${C5}╔═╗╦ ╦╔═╗╔═╗╔═╗╔╦╗╔═╗  ╔═╗╔═╗${NC}"
    echo -e "${C6}╠═╣║║║║╣ ╚═╗║ ║║║║║╣   ╠═╝║ ║${NC}"
    echo -e "${C7}╩ ╩╚╩╝╚═╝╚═╝╚═╝╩ ╩╚═╝  ╩  ╚═╝${NC}"
    echo -e "${C8}★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★${NC}"
    echo -e "${C9}Terminal Kubernetes context switcher${NC}"
    echo -e "${C10}With retro-wave UI design${NC}"
    echo ""
}

# Function to install kcs
install_kcs() {
    print_banner

    echo -e "${C5}Installing Kube Switch (kcs)...${NC}"
    echo ""

    # Create directory for binary
    mkdir -p "$BIN_DIR"

    # Download the code directly
    echo -e "${C6}Downloading code...${NC}"
    curl -L -s https://raw.githubusercontent.com/sky0ps/kube-switch/main/main.go -o "main.go"

    # Build the binary
    echo -e "${C7}Building binary...${NC}"
    go build -o "$INSTALL_PATH" main.go

    # Clean up
    rm main.go

    # Check if the directory is in PATH
    if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
        echo -e "${C8}Adding $BIN_DIR to PATH...${NC}"
        echo "export PATH=\"\$PATH:$BIN_DIR\"" >> "$HOME/.bashrc"
        echo -e "${C9}Please run 'source ~/.bashrc' to update your PATH${NC}"
    fi

    # Make the binary executable
    chmod +x "$INSTALL_PATH"

    # Final message
    echo -e "\n${C10}Kube Switch (kcs) has been installed successfully!${NC}"
    echo -e "${C9}Run 'kcs' to start the application.${NC}"
    echo -e "${C8}To uninstall, run 'kcs --uninstall' or this script with --uninstall${NC}"
}

# Function to uninstall kcs
uninstall_kcs() {
    print_banner
    
    echo -e "${C5}Uninstalling Kube Switch (kcs)...${NC}"
    
    # Check if binary exists
    if [ -f "$INSTALL_PATH" ]; then
        # Remove binary
        rm -f "$INSTALL_PATH"
        echo -e "${C7}Removed binary from $INSTALL_PATH${NC}"
        
        # Check if PATH was modified
        if grep -q "export PATH=\"\$PATH:$BIN_DIR\"" "$HOME/.bashrc"; then
            echo -e "${C8}NOTE: The PATH entry in ~/.bashrc was not removed.${NC}"
            echo -e "${C9}If you don't have other programs in $BIN_DIR, you may want to remove this line:${NC}"
            echo -e "${C10}    export PATH=\"\$PATH:$BIN_DIR\"${NC}"
        fi
        
        echo -e "\n${C10}Kube Switch (kcs) has been uninstalled successfully!${NC}"
    else
        echo -e "${C5}Kube Switch (kcs) does not appear to be installed.${NC}"
    fi
}

# Process command line arguments
if [ "$1" == "--uninstall" ]; then
    uninstall_kcs
else
    install_kcs
fi

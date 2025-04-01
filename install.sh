#!/bin/bash

# Simple Kube Switch (kcs) Installer

set -e # Exit immediately if a command exits with a non-zero status.

# Configuration
REPO_URL="https://github.com/sky0ps/kube-switch.git" # Using git clone is generally better
# Or keep raw file download if preferred, but cloning gets go.mod/go.sum if they exist
RAW_MAIN_GO_URL="https://raw.githubusercontent.com/sky0ps/kube-switch/main/main.go"
INSTALL_DIR="/usr/local/bin"
TMP_DIR=$(mktemp -d)

# Colors for output
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Cool Banner Function
print_banner() {
    echo -e "${BLUE}"
    echo "╔╗ ╔═╗  ╦╔═╦ ╦╔╗ ╔═╗  ╔═╗╦ ╦╦╔╦╗╔═╗╦ ╦"
    echo "╠╩╗║ ║  ╠╩╗║ ║╠╩╗║╣   ╚═╗║║║║ ║ ║  ╠═╣"
    echo "╚═╝╚═╝  ╩ ╩╚═╝╚═╝╚═╝  ╚═╝╚╩╝╩ ╩ ╚═╝╩ ╩"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "Terminal Kubernetes context switcher"
    echo "With retro-wave UI design"
    echo -e "${NC}"
}

# Cleanup function
cleanup() {
    echo -e "${YELLOW}Cleaning up temporary directory: $TMP_DIR${NC}"
    rm -rf "$TMP_DIR"
}

# Trap cleanup function to run on exit
trap cleanup EXIT

# --- Installation Steps ---

print_banner
echo -e "${GREEN}Installing Kube Switch (kcs)...${NC}\n"

# Check for Go installation
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed. Please install Go (version 1.16 or higher is recommended).${NC}"
    exit 1
fi

echo "Using temporary directory: $TMP_DIR"
cd "$TMP_DIR"

# --- Download Code ---
echo "Downloading code..."
if curl --fail -L "$RAW_MAIN_GO_URL" -o main.go; then
    echo "Downloaded main.go successfully."
else
    echo -e "${RED}Error: Failed to download main.go from $RAW_MAIN_GO_URL${NC}"
    exit 1
fi

# --- Setup Go Modules ---
echo "Setting up Go modules..."
# Initialize go.mod (since we only downloaded main.go)
go mod init kube-switch-temp > /dev/null 2>&1
echo "Running 'go mod tidy' to download dependencies and verify checksums..."
# <<< FIX HERE: Use 'go mod tidy' >>>
if go mod tidy; then
   echo "'go mod tidy' completed successfully."
else
   echo -e "${RED}Error: 'go mod tidy' failed. Check Go environment and network.${NC}"
   exit 1
fi
# Note: 'go get ...' is no longer strictly needed here as 'go mod tidy' handles fetching.

# --- Build ---
echo "Building binary..."
# Build the binary. -s -w strips debug symbols andDWARF info to make binary smaller.
if go build -ldflags="-s -w" -o kcs .; then
    echo "Build successful."
else
    echo -e "${RED}Error: Go build failed.${NC}"
    # You might want to add 'cat go.mod' or 'cat go.sum' here for debugging if it fails
    exit 1
fi


# --- Install ---
echo "Attempting to install kcs to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
    if mv kcs "$INSTALL_DIR/"; then
        echo -e "${GREEN}Successfully installed kcs to $INSTALL_DIR/kcs${NC}"
    else
        echo -e "${RED}Error: Failed to move kcs to $INSTALL_DIR. Try running with sudo.${NC}"
        exit 1
    fi
else
    echo -e "${YELLOW}Warning: No write permissions for $INSTALL_DIR. Attempting with sudo...${NC}"
    if sudo mv kcs "$INSTALL_DIR/"; then
         echo -e "${GREEN}Successfully installed kcs to $INSTALL_DIR/kcs (using sudo)${NC}"
    else
        echo -e "${RED}Error: Failed to install kcs with sudo. Please check permissions for $INSTALL_DIR.${NC}"
        exit 1
    fi
fi

echo -e "\n${GREEN}Installation Complete! You can now run 'kcs'.${NC}"

# Cleanup is handled by the trap

exit 0

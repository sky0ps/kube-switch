#!/bin/bash

set -e

# Colors
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# ASCII Art Banner
echo -e "${PURPLE}"
echo "  _  __     _            _____         _ _       _     "
echo " | |/ /    | |          / ____|       (_) |     | |    "
echo " | ' /_   _| |__   ___ | (_____      ___| |_ ___| |__  "
echo " |  <| | | | '_ \ / _ \ \___ \ \ /\ / / | __/ __| '_ \ "
echo " | . \ |_| | |_) |  __/ ____) \ V  V /| | || (__| | | |"
echo " |_|\_\__,_|_.__/ \___||_____/ \_/\_/ |_|\__\___|_| |_|"
echo -e "${NC}"

echo -e "${BLUE}Installing Kube Switch (kcs)...${NC}"
echo ""

# Create directory for binary
BIN_DIR="$HOME/bin"
mkdir -p "$BIN_DIR"

# Download the code directly
echo -e "${BLUE}Downloading code...${NC}"
curl -L -s https://raw.githubusercontent.com/sky0ps/kube-switch/main/main.go -o "main.go"

# Build the binary
echo -e "${BLUE}Building binary...${NC}"
go build -o "$BIN_DIR/kcs" main.go

# Clean up
rm main.go

# Check if the directory is in PATH
if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
  echo -e "${YELLOW}Adding $BIN_DIR to PATH...${NC}"
  echo "export PATH=\"\$PATH:$BIN_DIR\"" >> "$HOME/.bashrc"
  echo -e "${YELLOW}Please run 'source ~/.bashrc' to update your PATH${NC}"
fi

# Final message
echo -e "\n${GREEN}Kube Switch (kcs) has been installed successfully!${NC}"
echo -e "${BLUE}Run 'kcs' to start the application.${NC}"

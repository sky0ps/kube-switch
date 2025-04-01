# 🌌 Kube Switch (kcs)

```
██╗  ██╗██╗   ██╗██████╗ ███████╗    ███████╗██╗    ██╗██╗████████╗ ██████╗██╗  ██╗
██║ ██╔╝██║   ██║██╔══██╗██╔════╝    ██╔════╝██║    ██║██║╚══██╔══╝██╔════╝██║  ██║
█████╔╝ ██║   ██║██████╔╝█████╗      ███████╗██║ █╗ ██║██║   ██║   ██║     ███████║
██╔═██╗ ██║   ██║██╔══██╗██╔══╝      ╚════██║██║███╗██║██║   ██║   ██║     ██╔══██║
██║  ██╗╚██████╔╝██████╔╝███████╗    ███████║╚███╔███╔╝██║   ██║   ╚██████╗██║  ██║
╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝    ╚══════╝ ╚══╝╚══╝ ╚═╝   ╚═╝    ╚═════╝╚═╝  ╚═╝
```

A retro-themed terminal UI for switching between Kubernetes contexts and namespaces.

## ✨ Features

- 📋 **List and switch** between Kubernetes contexts with a stylish terminal UI
- 🔄 **Switch namespaces** within contexts
- 🎨 **Color-coded** contexts based on environment (production, staging, development)
- 🚨 **Safety confirmations** for production environments
- 💻 **Retro-wave inspired** color theme

## 🚀 Installation

### Option 1: Download binary

```bash
# Create directory for binary
mkdir -p ~/bin

# Download and install
curl -L https://raw.githubusercontent.com/sky0ps/kube-switch/main/install.sh | bash

# Add to PATH if needed
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Option 2: Build from source

```bash
# Clone the repository
git clone https://github.com/sky0ps/kube-switch.git
cd kube-switch

# Build the binary
go build -o kcs

# Install to your path
mkdir -p ~/bin
cp kcs ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

## 🎮 Usage

Simply run `kcs` in your terminal:

```bash
kcs
```

### Navigation

- **↑/↓** - Navigate through contexts
- **Enter** - Select a context
- **Esc** - Go back
- **Ctrl-C** - Quit

## 📝 License

This project is licensed under the MIT License.

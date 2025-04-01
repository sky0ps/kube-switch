# 🌌 Kube Switch (kcs)

```
 ╦╔═╦ ╦╔╗ ╔═╗  ╔═╗╦ ╦╦╔╦╗╔═╗╦ ╦
 ╠╩╗║ ║╠╩╗║╣   ╚═╗║║║║ ║ ║  ╠═╣
 ╩ ╩╚═╝╚═╝╚═╝  ╚═╝╚╩╝╩ ╩ ╚═╝╩ ╩
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

A retro-themed terminal UI for switching between Kubernetes contexts and namespaces, with a gorgeous retrowave color scheme.

##  Features

-  **List and switch** between Kubernetes contexts with a stylish terminal UI
-  **Switch namespaces** within contexts
-  **Color-coded** contexts based on environment (production: purple, staging: magenta, development: blue)
-  **Safety confirmations** for production environments
-  **Retro-wave inspired** color theme with purple and blue gradients

## 🚀 Installation

### Option 1: One-line installer

```bash
curl -L https://raw.githubusercontent.com/sky0ps/kube-switch/main/install.sh | bash
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

## 🧹 Uninstallation

To uninstall Kube Switch, you can run either:

```bash
# Option 1: Using the installer script with --uninstall flag
curl -L https://raw.githubusercontent.com/sky0ps/kube-switch/main/install.sh | bash -s -- --uninstall

# Option 2: Direct removal
rm ~/bin/kcs
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

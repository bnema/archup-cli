#!/bin/bash
# Extract Desktop Foundation packages from Omarchy
# This script analyzes Omarchy's 148-package monolith and extracts
# only the infrastructure packages needed for Tier 2 (Desktop Foundation)

set -e

OMARCHY_BASE="vendor/omarchy/install/omarchy-base.packages"
FOUNDATION_OUT="defaults/packages/desktop-foundation.packages"

echo "=== Extracting Desktop Foundation Packages from Omarchy ==="
echo ""

# Check if Omarchy submodule exists
if [ ! -f "$OMARCHY_BASE" ]; then
  echo "ERROR: Omarchy submodule not found at $OMARCHY_BASE"
  echo "Run: git submodule update --init --recursive"
  exit 1
fi

# Create output file
mkdir -p "$(dirname "$FOUNDATION_OUT")"

cat > "$FOUNDATION_OUT" << 'EOF'
# Desktop Foundation Layer (Tier 2)
# Extracted from Omarchy's 148-package base
# These are universal desktop requirements for ANY compositor
#
# NOTE: The following are already installed in Tier 1 (archup):
#   - iwd (WiFi)
#   - bluez, bluez-utils (Bluetooth base)
#   - bash-completion
#   - fzf, ripgrep, bat, eza, fd, starship, zoxide (modern CLI tools)
#   - git, neovim, man-db, sudo
#
# This list contains ONLY Tier 2 desktop infrastructure

# === GRAPHICS & HARDWARE ===
# Note: Hardware-specific drivers (mesa, vulkan) are in separate files
# and installed based on GPU detection

# === AUDIO STACK (PipeWire) ===
pipewire
wireplumber
# Note: pipewire-pulse, pipewire-alsa, pipewire-jack installed conditionally

# === WAYLAND CORE ===
qt5-wayland
qt6-wayland
xdg-desktop-portal-gtk
polkit-gnome

# === NETWORK SERVICES (mDNS + Avahi) ===
avahi
nss-mdns
# Note: iwd already installed in Tier 1

# === BLUETOOTH UI ===
blueberry
# Note: bluez/bluez-utils already in Tier 1

# === PRINTING (CUPS + mDNS discovery) ===
cups
cups-browsed
cups-filters
cups-pdf
system-config-printer

# === FONTS & RENDERING ===
fontconfig
noto-fonts
noto-fonts-cjk
noto-fonts-emoji
noto-fonts-extra
ttf-cascadia-mono-nerd
# Note: ttf-jetbrains-mono-nerd already in Tier 1

# === AUTHENTICATION & KEYRINGS ===
gnome-keyring
libsecret

# === FILE SYSTEMS (GVFS support) ===
gvfs-mtp
gvfs-nfs
gvfs-smb

# === POWER MANAGEMENT ===
power-profiles-daemon
brightnessctl

# === SESSION MANAGEMENT ===
uwsm
# Note: SDDM is optional, installed only if user selects display manager

# === ADDITIONAL TOOLS FROM OMARCHY ===
# Tools that enhance desktop experience
playerctl
imagemagick
less
man
plocate
tree-sitter-cli
unzip
wl-clipboard

# === AUR PACKAGES (Optional) ===
# These are from Omarchy and can be optionally installed
# impala          # WiFi TUI manager (user selects in wizard)
# wiremix         # Audio mixer TUI (user selects in wizard)

EOF

echo "✓ Desktop Foundation packages extracted to: $FOUNDATION_OUT"
echo ""

# Count packages
PACKAGE_COUNT=$(grep -v '^#' "$FOUNDATION_OUT" | grep -v '^$' | wc -l)
echo "  Total packages: $PACKAGE_COUNT"
echo ""

# Now create hardware-specific package files
echo "Creating hardware-specific package files..."
echo ""

# Intel packages
cat > "defaults/packages/hardware-intel.packages" << 'EOF'
# Intel GPU packages
mesa
lib32-mesa
vulkan-intel
lib32-vulkan-intel
intel-media-driver
libva-intel-driver
EOF

echo "✓ Created: defaults/packages/hardware-intel.packages"

# NVIDIA packages
cat > "defaults/packages/hardware-nvidia.packages" << 'EOF'
# NVIDIA GPU packages
nvidia-dkms
nvidia-utils
lib32-nvidia-utils
nvidia-settings
EOF

echo "✓ Created: defaults/packages/hardware-nvidia.packages"

# AMD packages
cat > "defaults/packages/hardware-amd.packages" << 'EOF'
# AMD GPU packages
mesa
lib32-mesa
vulkan-radeon
lib32-vulkan-radeon
libva-mesa-driver
lib32-libva-mesa-driver
EOF

echo "✓ Created: defaults/packages/hardware-amd.packages"

echo ""
echo "=== Extraction Complete ==="
echo ""
echo "Next steps:"
echo "  1. Review defaults/packages/desktop-foundation.packages"
echo "  2. Create compositor-specific package files (niri, hyprland, etc.)"
echo "  3. Implement hardware detection in internal/install/hardware.go"
echo ""

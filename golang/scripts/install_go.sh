#!/bin/bash

# Go Installer for WSL
set -e

echo "Go Installer for WSL"
echo "===================="

# Check if Go is already installed
if command -v go &> /dev/null; then
    echo "Go is already installed!"
    go version
    exit 0
fi

# Get latest Go version
echo "Fetching latest Go version..."
VERSION=$(curl -s -L https://golang.org/VERSION?m=text | head -n1 | tr -d '\n')
ARCHIVE_NAME="${VERSION}.linux-amd64.tar.gz"
DOWNLOAD_URL="https://golang.org/dl/${ARCHIVE_NAME}"

echo "Installing ${VERSION}..."

# Download Go to user's home directory first
echo "Downloading ${ARCHIVE_NAME}..."
cd ~
curl -LO "${DOWNLOAD_URL}"

# Remove any existing Go installation
sudo rm -rf /usr/local/go

# Extract Go
echo "Extracting Go..."
sudo tar -C /usr/local -xzf "${ARCHIVE_NAME}"

# Clean up
rm "${ARCHIVE_NAME}"

# Add to PATH
echo "Adding Go to PATH..."
echo '' >> ~/.profile
echo '# Go' >> ~/.profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile

# Create symlink for easier access
sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go

echo ""
echo "Go installation completed!"
echo "${VERSION} has been installed to /usr/local"
echo ""
echo "To use Go immediately, run:"
echo "  source ~/.profile"
echo ""
echo "Or restart your terminal."
echo ""
echo "Verify installation with:"
echo "  go version"
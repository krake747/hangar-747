#!/bin/bash

# Kubernetes Setup Script
# This script installs Minikube, kubectl, and K9s on Ubuntu/WSL

set -e  # Exit on any error

echo "Starting Kubernetes tools installation..."

# Update package list
echo "Updating package list..."
sudo apt update

# Install dependencies
echo "Installing dependencies..."
sudo apt install -y curl wget apt-transport-https

# Install Minikube
echo "Installing Minikube..."
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
rm minikube_latest_amd64.deb

# Install kubectl
echo "Installing kubectl..."
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
rm kubectl

# Install K9s
echo "Installing K9s..."
# Get latest version
K9S_VERSION=$(curl -s https://api.github.com/repos/derailed/k9s/releases/latest | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
curl -LO "https://github.com/derailed/k9s/releases/download/${K9S_VERSION}/k9s_Linux_amd64.tar.gz"
tar -xvzf k9s_Linux_amd64.tar.gz
sudo mv k9s /usr/local/bin/
rm k9s_Linux_amd64.tar.gz

# Verify installations
echo "Verifying installations..."
minikube version
kubectl version --client
k9s version

echo "Installation complete! You can now start Minikube with 'minikube start'"
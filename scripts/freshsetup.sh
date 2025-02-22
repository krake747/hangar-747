#!/bin/bash

set -e
set -u
set -o pipefail

echo "Updating system..."
sudo apt update && sudo apt upgrade -y

# -+- Installing common dependencies -+-
sudo apt install -y curl wget git unzip apt-transport ca-certificates gnupg


# -+- Installing programming languages -+-


## Install Python
echo "Installing Python..."
sudo apt install -y python3 python3-pip python3-venv


## Install .NET
echo "Installing .NET..."
sudo add-apt-repository ppa:dotnet/backports
sudo apt install -y dotnet-sdk-9.0


## Install Golang
echo "Installing Golang..."
GO_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n 1)
wget "https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz"
sudo tar -C /usr/local -xzf "${GO_VERSION}.linux-amd64.tar.gz"
rm "${GO_VERSION}.linux-amd64.tar.gz"
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc


## Install Node
echo "Installing node via NVM..."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash


# -+- Installing tools -+-


## Install Lazygit
echo "Installing Lazygit for managing Git workflows in the terminal..."
LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | \grep -Po '"tag_name": *"v\K[^"]*')
curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/download/v${LAZYGIT_VERSION}/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz"
tar xf lazygit.tar.gz lazygit
sudo install lazygit -D -t /usr/local/bin/

## Install Docker
echo "Installing Docker to build some images and run some containers..."
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin


## Install Kubernetes CLI (kubectl)
echo "Installing the Kubernetes CLI (kubectl)..."
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
sudo chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg # allow unprivileged APT programs to read this keyring

# This overwrites any existing configuration in /etc/apt/sources.list.d/kubernetes.list
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo chmod 644 /etc/apt/sources.list.d/kubernetes.list   # helps tools such as command-not-found to work correctly
sudo apt-get install -y kubectl


## Install Helm
echo "Installing Helm to navigate k8s with charts..."
curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
sudo apt-get install apt-transport-https --yes
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get install helm


## Install k9s TUI
echo "Installing k9s to manage a Kubernetes cluster in style..."
curl -LO https://github.com/derailed/k9s/releases/download/v0.40.3/k9s_Linux_amd64.tar.gz
tar -xvzf k9s_Linux_amd64.tar.gz
sudo mv k9s /usr/local/bin/
rm k9s_Linux_amd64.tar.gz


# -+- Installing programming languages -+-
echo "Fresh setup complete! Checking versions..."
python3 --version
dotnet --version
go version
lazygit --version
docker --version
kubectl --version
helm version


echo "Done! Restart the session for changes to take effect..."

# hangar-747

My homelab playground.

## Tech Stack

-   WSL with Ununtu distro (for now)
-   Docker Engine
-   Kubernetes via minikube

## Setup for WSL

### Docker

### Minikube

Minikube is local Kubernetes. All it requires is a Docker container.

[Minikube Docs](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download)

```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
```

Verfiy the installation via:

```shell
minikube start
```

In case the `kubectl` client and server don't match, we can update this via:

```shell
curl -LO "https://dl.k8s.io/release/v1.32.0/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```

### K9s

K9s is a TUI to manage a Kubernetes cluster.

[K9s Docs](https://k9scli.io/)

**Note:** You might need to download the latest version and modify the tarball url accordingly.

```shell
curl -LO https://github.com/derailed/k9s/releases/download/v0.40.3/k9s_Linux_amd64.tar.gz
tar -xvzf k9s_Linux_amd64.tar.gz
sudo mv k9s /usr/local/bin/
rm k9s_Linux_amd64.tar.gz
```

Verfiy the version via:

```shell
k9s version
```

### Bash

Add an alias `d` for docker and `k` for `kubectl` to the `.bashrc` file. Open it and add these three lines.

```shell
alias d='docker'
alias k='kubectl'
complete -o default -F __start_kubectl k
```

Inside the `dotfiles` directory we can have look of my `.bashrc` file which contains aliases for `git`, `docker` and `k8s`.
 
### Lazygit

Sometimes I want to have a better overview of my git commit history. I am exploring `lazygit` as an alternative.

```shell
LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | \grep -Po '"tag_name": *"v\K[^"]*')
curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/download/v${LAZYGIT_VERSION}/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz"
tar xf lazygit.tar.gz lazygit
sudo install lazygit -D -t /usr/local/bin/
```

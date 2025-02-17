# hangar-747

My homelab playground.

## Tech Stack

- WSL with Ununtu distro (for now)
- Docker Engine
- Kubernetes via minikube

## Setup for WSL

### Docker

### Minikube

Minikube is local Kubernetes. All it requires is a Docker container.

[Minikube Docs](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download)

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
```

Verfiy the installation via:

```bash
minikube start
```

### K9s

K9s is a TUI to manage a Kubernetes cluster.

[K9s Docs](https://k9scli.io/)

**Note:** You might need to download the latest version and modify the tarball url accordingly.

```bash
curl -LO https://github.com/derailed/k9s/releases/download/v0.40.3/k9s_Linux_amd64.tar.gz
tar -xvzf k9s_Linux_amd64.tar.gz
sudo mv k9s /usr/local/bin/
rm k9s_Linux_amd64.tar.gz
```

Verfiy the version via:

```bash
k9s version
```

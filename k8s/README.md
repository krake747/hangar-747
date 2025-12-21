# Kubernetes (k8s)

This directory contains guides and configurations for learning Kubernetes fundamentals. It includes
setup instructions for local development tools like Minikube and K9s, along with examples of
deploying applications to a local cluster. This helps me build a strong foundation in container
orchestration by providing hands-on experience with Kubernetes resources and deployment strategies.

## Installation

We recommend using the automated script from root folder for installation:

```shell
bash scripts/k8s_setup.sh
```

This script installs Minikube, kubectl, and K9s in one go.

## Getting started with Minikube

To learn Kubernetes, we start with a local cluster. Minikube gives us a simple Kubernetes setup on
our computer using Docker. It lets us build and test apps without cloud costs. This approach helps
me understand container orchestration fundamentals without the complexity of cloud deployments.

```bash
minikube start
```

For more details, check the
[Minikube Docs](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download).

## Getting started with K9s

Once we have a Kubernetes cluster running, we need an efficient way to interact with it. K9s
provides a terminal-based user interface (TUI) that makes managing cluster resources intuitive.
Instead of typing long kubectl commands, we can navigate pods, services, deployments, and logs
interactively, which speeds up our development and debugging process.

```bash
k9s
```

For more details, visit the [K9s Docs](https://k9scli.io/).

[Notes for k8s commands](k8s.md)

## Aliases

Add an alias `d` for docker and `k` for `kubectl` to the `.bashrc` file. Open it and add these three
lines.

```shell
alias d='docker'
alias k='kubectl'
complete -o default -F __start_kubectl k
```

## Essential K8s manifests

- `apply.sh`: A script to apply all Kubernetes manifests in the current directory.
- `hangar-namespace.yml`: A manifest to create a dedicated namespace for hangar-related resources.
- `hangar-ingress.yml`: A manifest to set up ingress rules for routing traffic to hangar services.

## Additional Resources

[Notes for kubectl commands](k8s.md)

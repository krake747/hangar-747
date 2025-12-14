# DevOps Journey via hangar-747

This document outlines the roadmap, goals, and topics I will cover to build a strong foundational understanding of DevOps and cloud-native engineering.
I already work as a full-stack developer (frontend + backend), and now I want to extend my skills into platform engineering, automation, and cloud operations.

## Goals

Build a practical foundation in DevOps principles and modern cloud infrastructure.
To achieve this, I will learn how to design, deploy, and operate cloud-native applications.

## Tech Stack

* [Go](https://go.dev/) and [.NET](https://dotnet.microsoft.com/en-us/) for building cloud tooling and microservices
* [OpenTofu](https://opentofu.org/)/[Terraform](https://developer.hashicorp.com/terraform) for infrastructure as code
* [Azure](https://azure.microsoft.com/en-us/get-started/azure-portal) as the cloud platform
* [Kubernetes (K8s)](https://kubernetes.io/) as the orchestration layer
* [Helm](https://helm.sh/) for application packaging and deployment

I will also establish consistent workflows for CI/CD, environment management, and observability.

## Roadmap

I am currently progressing through the foundational stages of DevOps with a focus on cloud-native development and Azure.
This section outlines where I am right now and the concrete goals I plan to achieve next.

### 1. Learn Go Basics

**Goal:** Build simple cloud-native utilities and services.

**Focus areas:**

* basic syntax, modules, and struct/method patterns
* concurrency fundamentals
* building CLI tools
* building small HTTP services

`/golang` directory contains my Go learning projects, and the `dotnet/` directory contains my .NET learning projects.

This phase builds the foundational programming skills for developing cloud-native applications and tooling.

### 2. Build CI/CD Pipelines with GitHub Actions

**Goal:** Automatically build and test Go and .NET projects.

**Focus areas:**

* build container images
* push images to Azure Container Registry
* prepare workflows for future Kubernetes and infrastructure automation

This phase establishes the automation backbone of the entire DevOps workflow. 

See the `.github` folder for CI/CD workflow files and `infra/` for deployment and infrastructure automation workflows.

### 3. Learn OpenTofu (Terraform) to Create Cloud Environments

**Goal:** Provision Azure resources using infrastructure as code and build reusable OpenTofu modules.

**Focus areas:**

* Azure Container Registry (ACR) for storing Docker images
* Azure Container Apps (ACA) environment as an initial lightweight hosting platform
* Azure Storage Account for Terraform state storage

This phase introduces infrastructure as code principles, enabling automated and version-controlled cloud resource management.

For detailed Terraform configurations, see [/infra/azure/README.md](./infra/azure/README.md).
and a live deployed resource [Web Api](https://krake-api-dev.mangoplant-23fd7721.westeurope.azurecontainerapps.io/)


### 4. Learn Kubernetes fundamentals

**Goal:** Understand how to deploy, scale, and manage containerized applications using Kubernetes.

TBD

### 5. Learn Helm for Application Packaging and Deployment

**Goal:** Package, configure, and deploy applications consistently across environments using Helm charts.

TBD

---

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

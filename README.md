# DevOps Journey via hangar-747

This document outlines the roadmap, goals, and topics I will cover to build a strong foundational
understanding of DevOps and cloud-native engineering. I already work as a full-stack developer
(frontend and backend), and now I want to extend my skills into platform engineering, automation,
and cloud operations.

## Goals

Build a practical foundation in DevOps principles and modern cloud infrastructure. To achieve this,
I will learn how to design, deploy, and operate cloud-native applications.

## Tech stack

- [Go](https://go.dev/) and [.NET](https://dotnet.microsoft.com/en-us/) for building cloud tooling
  and microservices
- [OpenTofu](https://opentofu.org/)/[Terraform](https://developer.hashicorp.com/terraform) for
  infrastructure as code
- [Azure](https://azure.microsoft.com/en-us/get-started/azure-portal) as the cloud platform
- [Docker](https://www.docker.com/) for containerization of applications
- [Kubernetes (K8s)](https://kubernetes.io/) as the orchestration layer
- [Helm](https://helm.sh/) for application packaging and deployment

I will also establish consistent workflows for CI/CD, environment management, and observability.

## Roadmap

I am currently progressing through the foundational stages of DevOps with a focus on cloud-native
development and Azure. This section outlines where I am right now and the concrete goals I plan to
achieve next.

### 1. Learn Go basics

**Goal:** Build simple cloud-native utilities and services.

**Focus areas:**

- basic syntax, modules, and struct/method patterns
- concurrency fundamentals
- building CLI tools
- building small HTTP services

`/golang` directory contains my Go learning projects, and the `/dotnet` directory contains my .NET
learning projects.

This phase builds the foundational programming skills for developing cloud-native applications and
tooling.

### 2. Build CI/CD pipelines with GitHub Actions

**Goal:** Automatically build and test Go and .NET projects.

**Focus areas:**

- build container images
- push images to Azure Container Registry
- prepare workflows for future Kubernetes and infrastructure automation

This phase establishes the automation backbone of the entire DevOps workflow.

See the `.github` folder for CI/CD workflow files and `infra/` for deployment and infrastructure
automation workflows.

### 3. Learn Docker for Containerization

**Goal:** Containerize Go and .NET applications for consistent deployment.

**Focus areas:**

- write Dockerfiles for Go and .NET apps
- build and run Docker images locally
- push images to Azure Container Registry (ACR)
- understand container networking and volumes
- local development with Docker Compose

This phase enables packaging applications into portable containers for deployment across
environments.

See the `dotnet/` directory for Dockerfiles and the `compose` orchestration and containerization
examples and `infra/` for ACR integration.

### 4. Learn OpenTofu (Terraform) to create cloud environments

**Goal:** Provision Azure resources using infrastructure as code and build reusable OpenTofu
modules.

**Focus areas:**

- Azure Container Registry (ACR) for storing Docker images
- Azure Container Apps (ACA) environment as an initial lightweight hosting platform
- Azure Storage Account for Terraform state storage

This phase introduces infrastructure as code principles, enabling automated and version-controlled
cloud resource management.

For detailed Terraform configurations, see [/infra/azure/README.md](./infra/azure/README.md). and a
live deployed resource
[Web Api](https://krake-api-dev.mangoplant-23fd7721.westeurope.azurecontainerapps.io/) (The first
response might be slow due the container's cold start)

### 5. Learn Kubernetes fundamentals

**Goal:** Learn the basics of deploying and exposing services using Kubernetes with Minikube.

**Focus areas:**

- deploy a simple web API to Minikube cluster
- create Kubernetes deployments and services
- expose services via LoadBalancer, NodePort, or Ingress
- configure basic networking and port forwarding
- test API accessibility from outside the cluster

For hands-on examples, see the `/k8s` directory containing Kubernetes manifests and deployment
scripts.

### 6. Learn Helm for Application Packaging and Deployment

**Goal:** Package, configure, and deploy applications consistently across environments using Helm
charts.

**Focus areas:**

- understand Helm concepts (charts, templates, values, releases)
- create a Helm chart for an existing Kubernetes application
- parameterize deployments using `values.yaml`
- manage environment-specific configuration (dev, local, prod)
- install, upgrade, and rollback releases
- structure charts for reuse and maintainability

For hands-on examples, see the `/helm` directory containing Helm charts and example values files
used to deploy applications to local Kubernetes clusters.

### 7. Deploy to a local Kubernetes cluster (mini PC)

**Goal:** Deploy and run applications on a local Kubernetes cluster hosted on a mini PC, simulating
a small production-like environment.

**Focus areas:**

- install and run a lightweight Kubernetes distribution (e.g. k3s)
- deploy containerized APIs to the cluster
- create Kubernetes deployments, services, and ingress resources
- expose services to the local network
- configure basic storage and resource limits
- verify application accessibility from other devices on the network

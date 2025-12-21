# Kubernetes (k8s)

This directory contains guides and configurations for learning Kubernetes fundamentals. It includes
setup instructions for local development tools like Minikube and K9s, along with examples of
deploying applications to a local cluster. This helps me build a strong foundation in container
orchestration by providing hands-on experience with Kubernetes resources and deployment strategies.

## Essential K8s manifests

- `hangar-namespace.yml`: A manifest to create a dedicated namespace for hangar-related resources.
- `skyops.yml`: A manifest to deploy the SkyOps outside facing web api.
- `skyhelp.yml`: A manifest to deploy the SkyHelp internal microservice.

I run in the `/dotnet` directory `bash build.sh` to create the docker images for SkyOps and SkyHelp.
Then I run `bash load.sh` to load the docker images to Minikube and `bash apply.sh` in this
directory to deploy the manifests to my local Minikube.

We either portforward the services or use Minikube's service command to access the SkyOps API.

```bash
minikube service skyops --url -n hangar
```

We can test the SkyOps API which internally calls the SkyHelp microservice like this (depends on the
tunnel port):

```bash
curl http://127.0.0.1:41589/redirect/skyhelp
```

## Reflections

As a fullstack developer, diving into Kubernetes is essential for my workflows. I need to be able to
build, configure and deploy the web applications I create.

For this project, I focused on understanding core Kubernetes concepts like namespaces, deployments ,
services and config maps. Deployments run app containers. They ensure apps are always running.
Services expose my deployments so other parts of your app can reach them. ConfigMap stores
configuration data separately from code, making it easier to manage settings without changing the
app itself.

- SkyOps is the external-facing API created as a `NodePort` service.
- SkyHelp is an internal microservice created as a `ClusterIP` service.

For internal communication Services get DNS names like `skyhelp.hangar.svc.cluster.local`. This
allows SkyOps to call SkyHelp securely within the cluster.

I mostly followed standard Kubernetes patterns to deploy the SkyOps and SkyHelp microservices in
this repo. Setting up ingress for external access while keeping SkyHelp internal-only was a good
lesson in service communication.

Adding namespace isolation and proper resource limits expands my knowledge on production-ready
deployments. The local development setup with Minikube and K9s made experimentation much easier.

## Additional Resources

[Notes for kubectl commands](k8s.md)

### Installation

We recommend using the automated script from root folder for installation:

```shell
bash scripts/k8s_setup.sh
```

This script installs Minikube, kubectl, and K9s in one go.

### Getting started with Minikube

To learn Kubernetes, we start with a local cluster. Minikube gives us a simple Kubernetes setup on
our computer using Docker. It lets us build and test apps without cloud costs. This approach helps
me understand container orchestration fundamentals without the complexity of cloud deployments.

```bash
minikube start
```

For more details, check the
[Minikube Docs](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download).

### Getting started with K9s

Once we have a Kubernetes cluster running, we need an efficient way to interact with it. K9s
provides a terminal-based user interface (TUI) that makes managing cluster resources intuitive.
Instead of typing long kubectl commands, we can navigate pods, services, deployments, and logs
interactively, which speeds up our development and debugging process.

```bash
k9s
```

For more details, visit the [K9s Docs](https://k9scli.io/).

[Notes for k8s commands](k8s.md)

### Aliases

Add an alias `d` for docker and `k` for `kubectl` to the `.bashrc` file. Open it and add these three
lines.

```shell
alias d='docker'
alias k='kubectl'
complete -o default -F __start_kubectl k
```

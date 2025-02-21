# Helm

Helm helps me manage Kubernetes applications.

A *Chart* is a Helm package. It contains all of the resource definitions necessary to run an application, tool, or
service inside of a Kubernetes cluster.

A *Repository* is the place where charts can be collected and shared. See [Artifacthub.io](https://artifacthub.io/)

A *Release* is an instance of a chart running in a Kubernetes cluster. One chart can often be installed many times
into the same cluster.

Helm installs *charts* into Kubernetes, creating a new *release* for each installation. And to find new charts, you
can search Helm chart *repositories*.

## Setup Helm

We first need a Kubernetes cluster that is up and running. For now I am using Minikube.

To install helm on WSL we can run:

```shell
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

To update Helm to latest version:

```shell
helm repo update
```

## Initialize a Helm Chart repository

```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
```

## Install a Helm chart

This will install the `nginx` chart

```shell
helm install my-nginx bitnami/nginx
```

## List Helm releases

```shell
helm list
```

## Uninstall a Helm chart

```shell
helm uninstall my-nginx
```

## Getting Started with Helm

To learn more about Helm and how to write Helm templates I follow the official
[Getting started](https://helm.sh/docs/chart_template_guide/getting_started/) guide.

In that guide we create a Helm chart from scratch.

# Hangar 747

## Gotenberg

One of the microservices that I would want to have is [Gotenberg](https://gotenberg.dev/) which will allow me to
convert numerous document formats. I want to use helm to set it up in my homelab.

I am using this [Helm chart](https://artifacthub.io/packages/helm/maikumori/gotenberg) from ArtifactHub. I can add it
to my repository.

```shell
helm repo add maikumori https://maikumori.github.io/helm-charts
helm repo update
```

To install it, I run the following command

```shell
helm install gotenberg --namespace hangar-747 maikumori/gotenberg
```

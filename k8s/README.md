# Kubernetes (k8s)

## Hangar 747

The first thing I want to create is a **namespace** called `hangar-747` to organise my resources. Then I want to create
the **deployment** which will manage and create the pods. Finally, I want to create the **service** which will expose
my deployment to network traffic.

### Namespace

What is a `k8s` namespace?

A **namespace** is just a logical grouping of things. In this case, services, pods, nodes, ingresses, etc live in a
namespace.

How to list `k8s` namespaces?

```shell
k get namespaces
```

Alternatively, this is the shorthand command

```shell
k get ns
```

**Note:** Running the `-h | less` command is useful to quickly reference and navigate the `k8s` docs.

How to create a namespace file?

```shell
k create ns hangar-747 --dry-run=client -o yaml > namespace.yaml
```

Here, we can remove the `creationTimestamp: null` property, since it can't be set manually

How to apply the namespace?

```shell
k apply -f namespace.yaml
```

How to get the current context and then set it to this namespace? (This way we avoid writing `-n` everywhere)

```shell
k config current-context
k config set-context --current --namespace=hangar-747
```

How to describe a namespace>

```shell
k describe ns hangar-747
```

How to delete a namespace?

```shell
k delete ns hangar-747
```

### Deployment

What is a `k8s` deployment?

A **deployment** manages pods. We create the deployment instead of managing pods ourselves.

How to get a deployment?

```shell
k get deployment
```

Alternatively, this is the shorthand command

```shell
k get deploy
```

How to create a deployment?

The following command will create a deployment called `my-nginx` using the official `nginx` image with 3 replicas.

```shell
k create deploy my-nginx --image=nginx --replicas=3
```

How to create a deployment from a local image?

Minikube has its own Docker daemon, so we just need to load the image from our local Docker environment into Minikube's Docker registry.

```shell
minikube image load skyops-api:latest
```

How do know if the image got loaded correctly?

This command will let us point the our local Docker environment to Minikube's Docker daemon and allow us to list the
images.

```shell
eval $(minikube -p minikube docker-env)
docker images | grep skyops
```

To revert the changes we can run

```shell
eval $(minikube -p minikube docker-env --unset)
```

Once we know it is loaded in Minikube we run our deployment

```shell
k create deploy skyops --image=skyops-api --dry-run=client -o yaml > deployment.yaml
```

**Note:** If `k8s` is trying to pull the image even though it is available locally in Minikube then we want to set
`imagePullPolicy: IfNotPresent` in the deployment file.

```yaml
spec:
    containers:
        - name: skyops-api
          image: skyops-api:latest
          imagePullPolicy: IfNotPresent # Avoid pulling the image from an external registry
          resources: {}
```

How to apply the deployment?

```shell
k apply -f deployment.yaml
```

How to delete a deployment?

```shell
k delete deploy hangar-747
```

How can we access the pod from outside of the cluster without a service?

```
k get pods
k port-forward pod/<entire pod-name> 7001:8080
```

This will forward the port to 7001 so we can access it locally.

### Service

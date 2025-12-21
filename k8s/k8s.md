# Hangar 747

The first thing I want to create is a _namespace_ called `hangar` to organise my resources. Then I
want to create the _deployment_ which will manage and create the pods. Finally, I want to create the
_service_ which will expose my deployment to network traffic.

## Namespace

### What is a `k8s` namespace?

A **namespace** is just a logical grouping of things. In this case, services, pods, nodes,
ingresses, etc live in a namespace.

### How to list `k8s` namespaces?

```bash
k get namespaces
k get ns
```

**Note:** Running the `-h | less` command is useful to quickly reference and navigate the `k8s`
docs.

### How to create the `hangar` namespace file?

```bash
k create ns hangar --dry-run=client -o yml > hangar-namespace.yml
```

### How to apply the namespace to our k8s cluster?

```bash
k apply -f hangar-namespace.yml
```

### How to get the current context and then set it to this namespace?

This way we avoid writing `-n` everywhere.

```bash
k config current-context
k config set-context --current --namespace=hangar
```

### How to describe a `namespace`?

```bash
k describe ns hangar
```

### How to delete a namespace?

```bash
k delete ns hangar
```

## Deployment

### What is a `k8s` deployment?

A **deployment** manages pods. We create the deployment instead of managing pods ourselves.

### How to get a deployment?

```bash
k get deployment
k get deploy
```

### How to create a deployment?

The following command will create a deployment called `my-nginx` using the official `nginx` image
with 3 replicas.

```bash
k create deploy my-nginx --image=nginx --replicas=3
```

### How to create a deployment from a local image?

Minikube has its own Docker daemon, so we just need to load the image from our local Docker
environment into Minikube's Docker registry.

```bash
minikube image load skyops:latest
```

### How do know if the image got loaded correctly?

This command will let us point the our local Docker environment to Minikube's Docker daemon and
allow us to list the images.

```bash
eval $(minikube -p minikube docker-env)
docker images | grep skyops
```

To revert the changes we can run

```bash
eval $(minikube -p minikube docker-env --unset)
```

Once we know it is loaded in Minikube we run our deployment

```bash
k create deploy skyops --image=skyops --dry-run=client -o yml > deployment.yml
```

**Note:** If `k8s` is trying to pull the image even though it is available locally in Minikube then
we want to set `imagePullPolicy: IfNotPresent` in the deployment file.

```yml
spec:
    containers:
        - name: skyops
          image: skyops:latest
          imagePullPolicy: IfNotPresent # Avoid pulling the image from an external registry
          resources: {}
```

How to apply the deployment?

```bash
k apply -f deployment.yml
```

How to delete a deployment?

```bash
k delete deploy hangar
```

How can we access the pod from outside of the cluster without a service?

```
k get pods
k port-forward pod/<entire pod-name> 7001:8080
```

This will forward the port to 7001 so we can access it locally.

### Service

A service is like a load balancer which offers a stable and fixed address (IP or DNS name) to access
a set of pods and route traffic to them.

We need a service because pods are ephemeral. We don't expect a pod to have a long lifespan. So when
pods get added or removed we can't keep track of the contantly changing IP addresses. This is where
`k8s` services come in handy.

There are four different types of `k8s` services:

1. `ClusterIP` - Default and internal access only
2. `NodePort` - Accessible from the outside the cluster via `<NodeIP>:<NodePort>`
3. `LoadBalancer` - External load balancer (often in cloud)
4. `ExternalName` - Maps a service to an external DNS name

How to get a service?

```bash
k get service
```

Alternatively, this is the shorthand command

```bash
k get svc
```

Or we need to additionally specify the namespace:

```bash
k get svc -n=default
```

How to create a service?

**ClusterIP**

This is the default type for a Kubernetes service.

```bash
k expose deploy skyops --type=ClusterIP --port=7000 --target-port=8080 --dry-run=client -o yml > clusterip-service.yml
```

This will expose port 7000 to external clients (within the cluster), and route the traffic to
container port 8080.

Inside another pod in the same cluster we could try to contact it via: `curl http://skyops:7000`.

We can use `k port-forward service/skyops 7000:7000` to expose it and test via
`curl http://localhost:7000`

**NodePort**

```bash
k expose deploy skyops --type=NodePort --port=7000 --target-port=8080 --dry-run=client -o yml > service.yml
```

- `port: 7000`: The port exposed by the service.
- `targetPort: 8080`: The container port to which the service will forward traffic.

How to apply a service?

```bash
k apply -f service.yml
```

We do the NodePort example here. Once it is applied we run

```bash
k get svc skyops -o wide
```

We should see something like this:

| NAME   | TYPE     | CLUSTER-IP    | EXTERNAL-IP | PORT(S)        | AGE   | SELECTOR   |
| ------ | -------- | ------------- | ----------- | -------------- | ----- | ---------- |
| skyops | NodePort | 10.102.167.86 | <none>      | 7000:30372/TCP | 5m31s | app=skyops |

This creates a tunnel to access (might be related to using docker driver on linux)

```bash
minikube service skyops -n hangar
```

This provides a url that we can access.

## ConfigMap

## Ingress

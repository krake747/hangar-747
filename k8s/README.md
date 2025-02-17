# Kubernetes (k8s)

## Hangar 747

### Namespace

The first thing I want to create is a namespace called `hangar-747`.

What exactly is a namespace?

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

**Note:** Running the `-h` command is useful to quickly reference the `k8s` docs.

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

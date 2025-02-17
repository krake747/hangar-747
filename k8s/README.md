# Kubernetes (k8s)

The first thing I want to create is a namespace called `hangar-747`.

## Hangar 747 Namespace

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

How to apply the namespace?

```shell
k apply -f namespace.yaml
```

How to get the current context and then set it to this namespace? (This way we avoid writing `-n` everywhere)

```shell
k config current-context
k config set-context --current --namespace=hangar-747
```

How to delete a namespace?

```shell
k delete ns hangar-747
```

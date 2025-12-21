# Helm Charts for Hangar

This directory contains Helm charts for deploying Hangar applications to Kubernetes.

## Reflections

As a fullstack developer, using Helm is important for managing complex Kubernetes deployments. It
allows me to package applications with configurations, dependencies, and templating for reusability
across environments. For now, there is no separate environment, but Helm charts will make
that transition easier in the future.

For this project, I focused on core Helm concepts like charts, templates, values, and releases.
Charts bundle Kubernetes manifests into versioned packages. Templates enable dynamic rendering with
Helm templating. Values provide parameterization for customization. Releases track installations and
upgrades.

- SkyOps chart packages the external API with NodePort service and environment configurations.
- SkyHelp chart packages the internal microservice with ClusterIP service and resource limits.

This modular approach with separate charts improved maintainability compared to raw YAMLs.
Integrating Helm into CI/CD pipelines ensures consistent, reproducible deployments.

## Available Charts

- `helm/skyops`: Chart for SkyOps (external API service)
- `helm/skyhelp`: Chart for SkyHelp (internal microservice)

## Installation

```bash
# Install SkyHelp first (dependency for SkyOps)
helm install skyhelp ./helm/skyhelp --create-namespace --namespace hangar

# Then install SkyOps
helm install skyops ./helm/skyops --namespace hangar
```

Or use the simple script: `bash apply.sh` (from the `helm/` directory)

## Dry Run (Validation)

Validate charts before deployment

### Dry-run install

```bash
helm install skyhelp ./helm/skyhelp --namespace hangar --dry-run=client
helm install skyops ./helm/skyops --namespace hangar --dry-run=client
```

### Dry-run templates

```bash
helm template skyhelp ./helm/skyhelp --namespace hangar
helm template skyops ./helm/skyops --namespace hangar
```

## Upgrading

```bash
helm upgrade skyhelp ./helm/skyhelp --namespace hangar
helm upgrade skyops ./helm/skyops --namespace hangar
```

## Uninstalling

```bash
helm uninstall skyops --namespace hangar
helm uninstall skyhelp --namespace hangar
```

## Configuration

Customize deployments by editing `values.yaml` in each chart directory. Common parameters:

- `image`: Docker image tag
- `replicas`: Number of pod replicas
- `service.port`: Service port
- `resources`: CPU/memory limits and requests

## Additional Resources

- [Helm Getting Started](https://helm.sh/docs/chart_template_guide/getting_started/)
- [Helm Chart Best Practices](https://helm.sh/docs/chart_best_practices/)
- [Helm Template Guide](https://helm.sh/docs/chart_template_guide/)

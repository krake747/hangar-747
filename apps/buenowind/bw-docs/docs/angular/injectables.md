# Injectables

In Angular, **injectables** such as services, route guards, and interceptors are injected using the
`inject()` function. **Standalone components** manage their **template context** explicitly through
the `imports: []` array, giving them direct control over their dependencies.

Unlike standalone components, injectables rely on the **injector hierarchy**, which determines how
services are provided and shared across the application. The `inject()` function is preferred over
**constructor injection** because the latter requires a workaround in `tsconfig.json`
to function correctly. Specifically, we must set:

```json
{
    "compilerOptions": {
        "useDefineForClassFields": false
    }
}
```

## Other injectors

There are other even more fine-grained injectors like ElementInjector
which is configured with the `providers: [ ]` array of a component or a
directive and allows us to create a unique instance of a given injectable for
each instance of component (or a directive) on which it was registered.

[Hierarchical injectors](https://angular.dev/guide/di/hierarchical-dependency-injection)

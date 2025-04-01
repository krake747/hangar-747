# Routing based lazy loading

In Angular, routing-based lazy loading is achieved through dynamic imports in the route configuration:

```ts
loadChildren: () => import('feature/<feature-name>/<feature-name>.routes')
```

This approach is also used for lazy-loaded sub-features, with the parent feature referencing child
routes using `loadChildren`. While `loadComponent` allows lazy loading of standalone components,
`loadChildren` is preferred for its flexibility, as it simplifies the addition of child routes.

Lazy bundles are extracted by the bundler when dynamic `import()` statements are encountered.
This lazy loading can be applied to both the main application and nested lazy features,
and the nesting can be as deep as needed.

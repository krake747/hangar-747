# Defer based lazy loading

Angular 17 introduced [`@defer`](https://angular.dev/guide/templates/defer), a simple way to lazily
load standalone components directly from the parent component's template. This is especially useful
for large components that significantly impact the bundle size, such as rich text editors, charts,
data tables, or maps.

While `@defer` is easy to use, it should not be applied to every component. Lazy loading
all components can lead to inefficient loading, especially in nested components where each one must
load sequentially, causing a "waterfall" of requests that degrades performance and user experience.

Instead, it's best to use `@defer` for heavy components that are either from
third-party libraries or are large local components. For most other components, eager loading
is more efficient, but for large, resource-heavy components, splitting them out into
lazy-loaded chunks can improve performance without delaying the entire feature.

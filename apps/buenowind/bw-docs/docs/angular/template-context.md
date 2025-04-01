# Template Context

The **template context** in standalone components refers to the collection of all components,
directives, and pipes required in a component's template. These dependencies must be explicitly
**imported** in the component's `imports: []` array.

With **NgModules**, the template context was managed at the module level. One key drawback was that
overly large template contexts limited the bundlers' ability to optimize lazy loading. Additionally,
the extra layer of indirection made it challenging for modern tooling that depends on single-file compilation.

With **standalone components**, every template dependency must be explicitly imported, resulting
in a clearer and more modular structure.

[The Most Important Thing You Need To Understand About Angular](https://angularexperts.io/blog/angular-template-context/)

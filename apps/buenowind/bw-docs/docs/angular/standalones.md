# Standalones

In older Angular versions, components, directives, and pipes had to be declared inside an `NgModule`
to be used.

With **standalone APIs**, we don't need `NgModules` anymore. Instead, we create components
with `standalone: true`.

If a standalone component needs to use another component, directive, or pipe, we import them
inside its `imports: []` array. This builds the **template context**, meaning it defines what
features the component can use in its template.

In Angular 19 the requirement to have `standalone: true` explicitly set, was removed. Standalone
is now the default when we scaffold a new project.

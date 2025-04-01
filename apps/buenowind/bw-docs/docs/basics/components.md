# Components

In Angular, **components** are the basic building blocks of an app.
They are defined using the `@Component` decorator, which includes a **selector** (the tag name) and
a **template** (HTML structure). Components can be nested within each other, where a parent component
can include a child component's selector in its template.

Components help keep an app modular by focusing on small, reusable tasks.
Communication between components happens through **inputs** (where a parent sends data to a child)
and **outputs** (where a child sends data to a parent).

```ts
@Component({
    selector: "rjn-root",
    imports: [RouterOutlet],
    template: `
        <router-outlet />
    `,
    styles: []
})
export class AppComponent {}
```

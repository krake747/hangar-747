# Single layout

Imagine an app where every page, whether it's the home page, orders page, dashboard, login, or signup,
shares the same basic structure. This means things like the header, sidebar, or footer stay the
same across all pages, while only the main content changes.

To achieve this, we create a **main layout** that wraps around everything.
This layout is placed inside the `AppComponent` (the root of the app) and includes a `<router-outlet>`.

So, whether the user is on the orders page, dashboard, or login page, they all fit inside this
single main layout, making the app look consistent and keeping the structure simple.

```ts title="app.component.ts"
@Component({
     standalone: true,
     selector: 'my-org-app',
     template: `<my-org-main-layout />`
})
export class AppComponent {}
```

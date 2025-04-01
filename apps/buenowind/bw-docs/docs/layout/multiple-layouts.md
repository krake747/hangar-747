# Multiple layout

Sometimes, an app needs **different layouts** for different parts of the app.

For example:

**Auth Layout** – Used for login and signup pages (without a sidebar or navigation).

**Main Layout** – Used after the user logs in, containing the header, sidebar, etc.

Instead of managing this in the `AppComponent`, we handle it in `app.routes.ts`, where we define
two main routing contexts:

- `''` (empty path) → Uses the `LoginLayoutComponent` (for authentication pages).
- `'/app'` → Uses the `MainLayoutComponent` (for all logged-in features).

Each layout acts as a wrapper and has its own `<router-outlet>` to load child pages.
This keeps everything organized, and the `AppComponent` only needs a single `<router-outlet>`.

The layouts themselves are placed in a `layout/` folder, keeping the structure clean and modular.

=== "app"

    ```ts title="app.component.ts"
    @Component({
        standalone: true,
        selector: 'my-org-app',
        template: `<router-outlet />`
    })
    export class AppComponent {}
    ```

=== "routes"

    ```ts title="app.routes.ts"
    export const routes: Routes = [
        // First layout
        {
            path: "",
            component: AuthLayoutComponent,
            children: [
                {
                    {
                        path: "login",
                        loadChildren: () => import("./feature/login/login.routes")
                    },
                    // signup feature, password recovery feature...
                }
            ]
        },
        // Second layout
        {
            path: "app",
            component: MainLayoutComponent,
            children: [
                {
                    {
                        path: "orders",
                        loadChildren: () => import("./feature/orders/orders.routes")
                    },
                    // dashboard, profile, settings feature...
                }
            ]
        },
    ];
    ```

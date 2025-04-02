# Nested lazy sub features

Sometimes, a feature is so big that it needs to be split into smaller parts (lazy sub-features).

- If these sub-features are **nested**, they appear alongside the parent (like a sidebar with sub-pages).
- If they are **separate**, they replace the parent view completely when opened.

These sub-features are stored inside the parent's folder and only load when needed,
keeping the app fast and organized.

=== "Appear in parent view"

    ```ts title="order.routes.ts"
    export const routes: Routes = [
        {
            path: "",
            // parent lazy feature component will always be displayed, needs <router-outlet />
            component: OrderComponent,
            children: [
                {
                    // will be displayed when navigated to and will be added to the parent view
                    path: "dashboard",
                    loadChildren: () => import("./dashboard/dashboard.routes")
                },
                {
                    path: "definitions",
                    loadChildren: () => import("./definitions/definitions.routes")
                },
            ]
        },
    ];
    ```

=== "Replace parent view"

    ```ts title="order.routes.ts"
    export const routes: Routes = [
        {
            path: "",
            component: OrderComponent, // parent lazy feature component
        }
        {
            // will be displayed when navigated to and will REPLACE the parent view
            path: "dashboard",
            loadChildren: () => import("./dashboard/dashboard.routes")
        },
        {
            path: "definitions",
            loadChildren: () => import("./definitions/definitions.routes")
        },
    ];
    ```

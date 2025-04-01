# Custom layout per feature

In some cases, different parts of an application are so distinct that each one requires its own unique
layout.

For example:

- The dashboard might require a layout that emphasizes graphs and widgets.
- The orders page might need a layout designed around tables and data entry.
- The admin panel might have a completely different structure with specific controls.

In this scenario:

- The `AppComponent` contains only a single `<router-outlet>`.
- The `layout/` folder will be empty, as each feature manages its own layout independently.
- The application-level route configuration remains simple, only defining the lazy-loaded features.

Each feature includes its own layout within its respective folder, such as `feature/dashboard/layout/`.
This approach allows every feature to define and customize its layout as needed without impacting
other parts of the application.

=== "app"

    ```ts title="app.component.ts"
    @Component({
        standalone: true,
        selector: 'my-org-app',
        template: `<router-outlet />`
    })
    export class AppComponent {}
    ```

=== "app.routes"

    ```ts title="app.routes.ts"
    export const routes: Routes = [
        {
            path: 'orders',
            loadChildren: () => import('./feature/orders/orders.routes')
        },
        // dashboard, profile, settings feature...
    ]
    ```

=== "dashboard.routes"

    ```ts title="dashboard.routes.ts"
    export const routes: Routes = [
        {
            path: '',
            component: DashboardLayoutComponent, // custom (per feature) layout
            children: [
                {
                    path: '',
                    component: DashboardComponent,
                },
                // other dashboard related routes...
                // other dashboard related sub lazy features...
            ]
        },
    ];
    ```

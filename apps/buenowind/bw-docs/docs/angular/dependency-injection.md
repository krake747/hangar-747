# Dependency Injection

Dependency Injection (DI) helps us manage how our app creates and shares objects.
Instead of manually creating new instances every time we need something, DI handles it for us.

**Easier Testing** - We can swap real dependencies with fake ones for testing.

**Shared Instances** - Different parts of our app can share a single instance of a service
(or use separate instances if needed).

**Less Hassle Creating Objects** - We don't need to manually pass all dependencies every time.

With DI, we can let Angular handle dependency resolution.
For example, in our `AppComponent`, we simply declare `MyService` in the constructor, and Angular
automatically provides the required dependencies:

```ts
export class AppComponent {
    myService = inject(MyService) {}
}
```

Without DI, we would need to manually create an instance of MyService along with all its dependencies:

```ts
export class AppComponent {
    myService = new MyService(new HttpClient(), new SomeOtherService());
}
```

In short, DI saves us time and makes our app easier to maintain.

## How Does Dependency Injection Work in Angular?

To use Dependency Injection (DI), we just add a parameter to our constructor and specify its type:

```ts
export class AppComponent {
    myService = inject(MyService) {}
}
```

This makes it seem like Angular automatically finds `MyService` and creates an instance of it,
but that's not exactly what happens. Instead, Angular treats `MyService` as a **token**,
which determines how and where an instance is provided.

### Providing Tokens

The most common way to provide a service is by using the `providedIn: 'root'` option:

```ts
@Injectable({
    providedIn: 'root'
})
export class MyService {}
```

This ensures that `MyService` is created once and shared across our entire application.

However, if we want a specific component to have its **own** instance of `MyService`,
we can override the default behavior using the `providers` array:

```ts
@Component({
    selector: 'app-my-component',
    template: `<p>hello</p>`,
    providers: [MyService]
})
export class MyComponent {}
```

Now, `MyComponent` will get a **separate** instance of `MyService`, instead of sharing the global one.

We can also provide services at the route level using the providers array inside route definitions.
This allows us to create service instances that are only available within a specific route and its
child components.

```ts
const routes: Routes = [
    {
        path: 'dashboard',
        component: DashboardComponent,
        providers: [DashboardService]
    }
];
```

When navigating to `/dashboard`, Angular will create a new instance of `DashboardService` that is only
available to `DashboardComponent` and its children.

If we navigate away and then return to `/dashboard`, a new instance of `DashboardService` will be created.

Using `providedIn: 'root'` is usually preferred because it supports **tree shaking**,
which removes unused code and keeps our app optimized. We prefer to use the route level definition if
we want to keep isolate the service to a feature.

### Overriding Tokens

Sometimes, we might want to **replace** a service with a different implementation.
We can do this by overriding the provider:

```ts
providers: [{ provide: MyService, useClass: CoolService }]
```

Now, whenever `MyService` is injected, Angular will actually use `CoolService` instead.
This technique is mostly useful for testing or when we need to swap out services dynamically.

## Summary

We use `providedIn: 'root'` for a single, shared instance of a service across our app.
We use the `providers` array inside a component if we want it to have its **own** instance of the service.
We can **override** tokens to replace a service with a different one, useful for testing or custom configurations.

# Function based logic

Angular now supports creating lightweight **function-based** route guards and HTTP interceptors without needing to write full class-based services with the `@Injectable` decorator.

- **Lightweight**: We can create function-based guards or interceptors as simple functions (e.g., a method of a route guard class).
- **Injection**: These functions can use `inject()` to inject other services, just like regular injectables.

```ts
export function roleBasedAuthGuardFactory(role: Role): CanActivateFn {
    return () => {
        const authService: AuthService = inject(AuthService);

        if (authService.isAuthenticated() && authService.hasRole(role)) {
            return true;
        } else {
            return router.createUrlTree(['/unauthorized']);
        }
    };
}
```

# Handling manual subscriptions

Sometimes, we need to manually subscribe to observables, like when sending data to a server or logging
analytics. Manual subscriptions are not ideal, but sometimes necessary.

If we do subscribe manually, we must unsubscribe to prevent memory leaks.
Even with single-value streams (like HTTP requests), not unsubscribing can cause issues if the component
is destroyed before the request finishes.

To unsubscribe safely, we can use the `takeUntilDestroyed` operator, which automatically unsubscribes
when the component is destroyed.

```ts
http = inject(HttpClient);
myData$ = this.http.get('someapi').pipe(takeUntilDestroyed());
```

Another option is the `takeUntil` operator and a `destroy$ = new Subject()`, which uses a notifier
to trigger unsubscription and the in the `ngDestroy` lifecycle hook.

```ts
ngOnDestroy(){
    this.destroy$.next();
    this.destroy$.complete();
}
```

In short, it's safest to always unsubscribe from observables, and `takeUntilDestroyed` is the
easiest and most reliable way to manage this.

# Higher order observables

A higher-order function is a function that either takes another function as input or returns one.
Similarly, a higher-order observable is an observable that emits other observables instead of plain
values like numbers or strings.

Most of the time in Angular, we work with observables that emit basic data.
But let's say we have a list of API URLs and want to fetch data from each one.
If we use `map()` with `http.get()`, we end up with a stream of observables, because each HTTP call
returns an observable. This means we now have a higher-order observable: a stream of streams.

Trying to handle this with nested `subscribe()` calls might seem like a solution, but it leads to
messy, non-reactive code. Instead, we should use a special operator like `mergeMap` or `switchMap`,
which can automatically flatten the stream of observables and give us the actual data.

When we end up with a stream of observables, we need to flatten it using the right operator
so we can work with the real values cleanly and reactively.

=== "Avoid nested subscriptions"

    ```ts
    const results: any[] = [];

    myObservable$.pipe(
        map(url => this.http.get(url))
    ).subscribe((result) => {
        result.subscribe((data) => results.push(data))
    })
    ```

=== "Better solution"

    ```ts
    const results: any[] = [];

    myObservable$.pipe(
        switchMap((url) => this.http.get(url))
    ).subscribe((data) => {
        results.push(data);
    });
    ```

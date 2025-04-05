# Flattening operators

Flattening operators help us deal with observables that emit other observables.
Instead of manually subscribing to each inner observable, we use a flattening operator to do it for us
It flattens everything down into one stream and gives us just the final values.
This way, we write cleaner code and stay reactive without nested subscriptions.

!!! summary

    1. We use `switchMap` if we only need the result from the **latest inner observable**.
    2. We use `concatMap` when we need the results from **all inner observables in order**.
    3. We use `mergeMap` when want **all results as quick as possible, disregarding order**.
    4. We use `exhaustMap` when we only care about the **current inner observable**.

## `switchMap`

When we use `switchMap`, we take a value from one stream and switch to a new inner observable,
like an HTTP request. If a new value comes in before the last one finishes, it cancels the old request
and starts a new one.

!!! example

    ```ts
    searchTerm$ = new Subject<string>();

    results$ = this.searchTerm$.pipe(
        debounceTime(300)
        switchMap(term => this.http.get(`/api/search?q=${term}`))
    );
    ```

    In this example we're creating a search feature where users can type into a search box,
    and we want to show live results from a server.
    But we want it to be smart, so if the user types quickly, we cancel the old search and only show
    the latest result.

    1. We receive a search term value from the **outer observable** (`searchTerm$`)
    2. We pass the searchterm into the get call, and subscribes to it (**inner observable**)
    3. Emits the value from the inner observable

## `concatMap`

`concatMap` is an operator that handles each value emitted by an observable one at a time.
It waits for each inner observable to complete before moving on to the next one.
This ensures that the results are emitted in the same order as the values were received,
but it can be slower because each request is handled sequentially, one after the other.

!!! example

    ```ts
    fileQueue$ = new Subject<File>();

    uploadResults$ = this.fileQueue$.pipe(
        concatMap(file => this.http.post('/api/upload', file))
    );
    ```

    In this example, we're simulating a scenario where we want to upload files one at a time in sequence.

    1. We receive a file from the **outer observable** (`fileQueue$`).
    2. We call the upload endpoint for the file, which triggers an HTTP request and subscribes to the
       **inner observable**.
    3. Once the first file upload completes, the next file is uploaded, and so on.

## `mergeMap`

We use `mergeMap` when we want to subscribe to all inner observables as soon as we receive new values,
without waiting for the previous ones to complete. Unlike `concatMap`, which processes one value at
a time in order, `mergeMap` processes them concurrently and emits results as they come in,
regardless of the order.

!!! example

    ```ts
    myObservable$.pipe(
        mergeMap(url => this.http.get(url))
    ).subscribe(val => console.log(val));
    ```

    If we're making several HTTP requests, `mergeMap` will handle all of them at once and return
    results as soon as they finish, without waiting for one to complete before starting the next.

    Here, if four requests are sent, `mergeMap` will process them in parallel and emit results as
    they complete. This is useful when we don’t care about the order of results and just want all
    data as quickly as possible.

## `exhaustMap`

The `exhaustMap` operator is used when we want to ignore new values from the outer observable while
there's an active inner observable. Once the inner observable completes, it will process the next
available value from the outer observable, but it won't process any values that arrive during the
ongoing inner observable.

!!! example

    Imagine you're at a bank, and someone else is already being served.
    If a new customer arrives while the first one is being served, the new customer will have to leave,
    they won't be helped until the first customer is done. This is how `exhaustMap` works.

    ```ts
    myObservable$.pipe(
        exhaustMap(url => this.http.get(url))
    ).subscribe(val => console.log(val));
    ```

    Here, if multiple requests are triggered, `exhaustMap` will ignore all new requests until the
    current request finishes. It's useful when we don’t want to start processing new values until
    the current one is fully handled.

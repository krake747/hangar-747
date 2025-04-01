---
tags:
    - experimental
---

# `resource`

The `resource()` API allows us to handle asynchronous operations with signals in a declarative way.
This is especially useful when making API requests and updating signals with the results once the
request completes, without needing to rely on imperative code.

Let’s say we want to fetch the player's game statistics from an API based on the current page the
player is on:

```ts
import { Component } from '@angular/core';

@Component({
    selector: 'app-game',
    template: `<p>Page: {{ currentPage() }}</p><p>Stats: {{ gameStats.value() }}</p>`
})
export class GameComponent {
    currentPage = signal(1);  // Current page number
    gameStats = resource({
        request: this.currentPage,  // Fetch data based on current page
        loader: ({ request, abortSignal }) => {
            return fetch(`https://example.com/game-stats/${request}/`, {
                signal: abortSignal,
            }).then((res) => res.json());
        },
    });
}
```

In this example:

- We use the `currentPage` signal to trigger an API request whenever the page changes.
- The `gameStats` resource will automatically update with the fetched data once the request finishes.
- The `abortSignal` ensures that if the page changes before the request finishes, the ongoing request
  is canceled.

We can also access additional status signals to know if the request is still loading or if there
was an error:

```typescript
this.gameStats.status();  // Check if the resource is still loading
this.gameStats.error();   // Check for any errors during the fetch
```

The `resource()` API simplifies handling asynchronous operations in Angular by making API requests declarative.
It allows us to automatically update signals with the results of the request, track the status of the
operation, and handle errors.

The `resource()` API also supports using observables, making it a flexible tool for dealing with
asynchronous operations. By using `resource()`, we can keep our code clean, declarative,
and efficient without relying on imperative effects.

## Links

[Improve the user experience of your application using (rx)resource](https://timdeschryver.dev/blog/improve-the-user-experience-of-your-application-using-rxresource)

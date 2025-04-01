---
tags:
    - experimental
---

# `rxResource`

While the `resource()` API supports handling asynchronous operations with the `fetch` API, it's more
common in Angular to use the `HttpClient` with observables for HTTP requests.
The `rxResource()` API is a variant that allows us to integrate the power of observables and `HttpClient`
with signals.

In this example, we will use `rxResource()` to fetch game data using `HttpClient` with observables:

```typescript
import { Component, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { rxResource } from "@angular/core/rxjs-interop";

@Component({
  selector: 'app-game',
  template: `<p>Page: {{ currentPage() }}</p><p>Stats: {{ gameStats.value() }}</p>`
})
export class GameComponent {
  http = inject(HttpClient);  // Injecting HttpClient

  currentPage = signal(1);  // Current page number
  gameStats = rxResource({
    request: this.currentPage,  // Fetch data based on the current page
    loader: (request) => this.http.get(`https://example.com/game-stats/${request}/`)
  });
}
```

- `rxResource()` integrates with `HttpClient` to fetch data declaratively using observables.
- `loader` function in `rxResource()` returns an observable (in this case, using `HttpClient`'s `get()`
  method).
- `request` is the signal (`currentPage`) that triggers the HTTP request when its value changes.
- The result of the HTTP request is automatically updated in the `gameStats` signal once the request
  completes.

Using `rxResource()` with `HttpClient` is a great way to handle asynchronous HTTP requests in Angular.
This approach leverages observables and integrates seamlessly with Angular’s reactive
architecture, making it a more natural choice for asynchronous operations in Angular applications.

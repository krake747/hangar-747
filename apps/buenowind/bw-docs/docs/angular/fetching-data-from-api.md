# Fetching data from Api

We set up the Angular `HttpClient` by using `provideHttpClient()` in `app.config.ts`.
We can then the `HttpClient` service to fetch the data from the Cat Fact API.

```ts
import { HttpClient } from "@angular/common/http";
import { ChangeDetectionStrategy, Component, inject } from "@angular/core";
import { toSignal } from "@angular/core/rxjs-interop";
import { Subject } from "rxjs";
import { map, startWith, switchMap } from "rxjs/operators";

type CatFact = {
    fact: string;
    length: number;
};

@Component({
    selector: "bw-cat-fact",
    imports: [],
    template: `
        <p class="p-2">{{ catFact() }}</p>
        <button
            class="cursor-pointer rounded-lg bg-sky-600 p-2 text-white hover:bg-sky-700"
            (click)="clicked$.next($event)"
        >
            Get New Fact
        </button>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class CatFactComponent {
    http = inject(HttpClient);

    clicked$ = new Subject<unknown>();

    catFact$ = this.clicked$.pipe(
        startWith({}),
        switchMap(() =>
            this.http.get<CatFact>("https://catfact.ninja/fact").pipe(map((c) => c.fact))
        )
    );

    catFact = toSignal(this.catFact$);
}
```

The API URL `https://catfact.ninja/fact` returns a random cat fact in JSON format.
The `HttpClient.get()` method is used to fetch the data from the API.
The `toSignal` function automatically subscribes to the observable and displays the data in the template.
The cat fact is displayed in a `<p>` tag, and there's a button that allows the user to fetch a new
random fact.
On page load, an HTTP request fetches a random cat fact and stores it in the signal.
When the user clicks the "Get New Fact" button, the component fetches and updates the displayed cat fact.

We can fetch data with the experimental signals using [`resource()`](../signal-apis/resource.md)
or [`rxResource`](../signal-apis/rx-resource.md). However, it is far more common currently to use `HttpClient`.

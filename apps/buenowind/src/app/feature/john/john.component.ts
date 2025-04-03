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
    selector: "bw-john",
    imports: [],
    template: `
        <p class="p-2 pb-3">{{ catFact() }}</p>
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
export class JohnComponent {
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

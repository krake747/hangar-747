import { AsyncPipe } from "@angular/common";
import { ChangeDetectionStrategy, Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { BehaviorSubject, combineLatest, concatMap, delay, from, map, of } from "rxjs";
import { NiceComponent } from "./nice/nice.component";

@Component({
    selector: "bw-home",
    imports: [RouterOutlet, NiceComponent, AsyncPipe],
    template: `
        <h1>Welcome to {{ title }}!</h1>
        <h1 class="text-3xl font-bold underline">Hello BuenoWind!</h1>
        @if (userMessage$ | async; as message) {
            <p>{{ message }}</p>
        }
        @if (currentPhoto$ | async; as photo) {
            <div>
                <img [src]="photo" alt="Current photo" />
            </div>
        }
        @if (userInfo$ | async; as userInfo) {
            <div>
                <p>Name: {{ userInfo.name }}</p>
                <p>Lucky Number: {{ userInfo.luckyNumber }}</p>
                <p>Attempts: {{ userInfo.attempts }}</p>
                <button
                    class="rounded bg-blue-500 p-1 text-white"
                    (click)="this.attempts$.next(this.attempts$.value + 1)"
                >
                    Attempt
                </button>
            </div>
        }
        <bw-nice />
        <router-outlet />
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class HomeComponent {
    title = "home";
    userMessage$ = of("Loading user data...");

    photos = ["photo1.jpg", "photo2.jpg", "photo3.jpg"];

    currentPhoto$ = from(this.photos).pipe(
        concatMap((photo) =>
            of(photo).pipe(
                delay(1000) // Delay for 1 second before showing each photo
            )
        )
    );

    name$ = of("John Doe");
    luckyNumber$ = of(7);
    attempts$ = new BehaviorSubject(0);

    userInfo$ = combineLatest([this.name$, this.luckyNumber$, this.attempts$]).pipe(
        map(([name, luckyNumber, attempts]) => ({ name, luckyNumber, attempts }))
    );
}

# Creation operators

In Angular, creation operators allow us to create completely new observable streams,
which are useful for emitting values or creating dynamic data flows.

Creation operators like `of` and `from` help us turn normal data into streams, and `combineLatest`
helps us merge multiple streams into one so we can handle all of them together efficiently.
This is particularly useful when working with multiple data points that need to update in sync,
such as in an Angular template using the `async` pipe.

## `of`

The `of` operator is used when we want to create an observable from a single value or a fixed set of
values.

**Use Case**: Displaying a default message or value before any actual data is fetched,
like showing a loading message before data arrives from an API.

```ts
import { of } from 'rxjs';

export class UserProfileComponent {
    userMessage$ = of('Loading user data...');
}
```

Here, we create an observable `userMessage$` that immediately emits the string `'Loading user data...'`.
This can be used in the template to display a loading message while waiting for actual data from an API.

```html
@if (userMessage$ | async; as message) {
    <p>{{ message }}</p>
}
```

[`of` marbles](https://rxjsmarbles.dev/of)

## `from`

The `from` operator is used when we want to convert a non-observable data type (like an array or a promise)
into an observable.

**Use Case**: Displaying a list of photo URLs one by one with a delay, as part of a slideshow feature.

```typescript
import { Component } from '@angular/core';
import { from, of } from 'rxjs';
import { concatMap, delay } from 'rxjs/operators';

@Component({
    selector: 'app-photo-slideshow',
    templateUrl: './photo-slideshow.component.html',
    styleUrls: ['./photo-slideshow.component.css']
})
export class PhotoSlideshowComponent {
    photos = ['photo1.jpg', 'photo2.jpg', 'photo3.jpg'];

    currentPhoto$ = from(this.photos).pipe(
        concatMap((photo) =>
            of(photo).pipe(
                delay(1000) // Delay for 1 second before showing each photo
            )
        )
    );
}
```

**Use Case**: We convert the array of photo URLs into an observable using `from`.
Then, using `concatMap`, we emit each photo one by one, with a 1-second delay between each emission
to simulate a slideshow effect.

```html
@if (currentPhoto$ | async; as photo) {
    <div>
        <img [src]="photo" alt="Current photo" />
    </div>
}
```

## `combineLatest`

The `combineLatest` operator is used when we have multiple observables and want to combine
their latest emitted values into one.

**Use Case**: Combining a user's name, lucky number, and the number of attempts they've made to
complete a task into a single observable stream for display.

```ts
import { Component } from '@angular/core';
import { BehaviorSubject, of } from 'rxjs';
import { combineLatest, map } from 'rxjs/operators';

@Component({
    selector: 'app-user-dashboard',
    templateUrl: './user-dashboard.component.html',
    styleUrls: ['./user-dashboard.component.css']
})
export class UserDashboardComponent {
    name$ = of("John Doe");
    luckyNumber$ = of(7);
    attempts$ = new BehaviorSubject(0);

    userInfo$ = combineLatest([this.name$, this.luckyNumber$, this.attempts$]).pipe(
        map(([name, luckyNumber, attempts]) => ({ name, luckyNumber, attempts }))
    );
}
```

**Use Case**: The `combineLatest` operator combines the latest values from `name$`, `luckyNumber$`,
and `attempts$` into a single observable.
We map the combined values into an object to easily display them in the template.

```html
@if (userInfo$ | async; as userInfo) {
    <div>
        <p>Name: {{ userInfo.name }}</p>
        <p>Lucky Number: {{ userInfo.luckyNumber }}</p>
        <p>Attempts: {{ userInfo.attempts }}</p>
        <button class (click)="this.attempts$.next(this.attempts$.value + 1)">Attempt</button>
    </div>
}
```

[`combineLatest` marbles](https://rxjsmarbles.dev/combineLatest)

# Commonly used operators

Here are simple example use cases for each operator in an Angular context:

## `map`

**Use Case**: Transforming a list of users, extracting only their names from an API response.

```typescript
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

constructor(private http: HttpClient) {}

getUserNames() {
    return this.http.get<User[]>('api/users').pipe(
        map(users => users.map(user => user.name)) // Transform user objects to just names
    );
}
```

[`map` marbles](https://rxjsmarbles.dev/map)

## `filter`

**Use Case**: Filtering a list of products to show only available ones.

```ts
import { of } from 'rxjs';
import { filter } from 'rxjs/operators';

const products = [
  { id: 1, name: 'Product A', available: true },
  { id: 2, name: 'Product B', available: false },
  { id: 3, name: 'Product C', available: true }
];

of(products).pipe(
    map(productList => productList.filter(product => product.available)) // Only show available products
).subscribe(filteredProducts => console.log(filteredProducts));
```

[`filter` marbles](https://rxjsmarbles.dev/filter)

## `tap`

**Use Case**: Logging form value changes for debugging.

```ts
import { FormControl } from '@angular/forms';
import { tap } from 'rxjs/operators';

const formControl = new FormControl();

formControl.valueChanges.pipe(
    tap(value => console.log('Form value changed:', value)) // Log each form change
).subscribe();
```

## `startWith`

**Use Case**: Setting a default value for a form control before the user interacts.

```ts
import { FormControl } from '@angular/forms';
import { startWith } from 'rxjs/operators';

const formControl = new FormControl();

formControl.valueChanges.pipe(
    startWith('default value') // Set initial default value
).subscribe(value => console.log('Current value:', value));
```

[`startWith` marbles](https://rxjsmarbles.dev/startWith)

## `distinctUntilChanged`

**Use Case**: Avoiding redundant HTTP requests when the search term in a form doesn’t change.

```ts
import { FormControl } from '@angular/forms';
import { distinctUntilChanged } from 'rxjs/operators';

const formControl = new FormControl();

formControl.valueChanges.pipe(
    distinctUntilChanged() // Emit value only if it changes
).subscribe(searchTerm => {
    console.log('Fetching data for:', searchTerm);
});
```

[`distinctUntilChanged` marbles](https://rxjsmarbles.dev/distinctUntilChanged)

## `debounceTime`

**Use Case**: Preventing a flood of HTTP requests while typing in a search bar.

```ts
import { FormControl } from '@angular/forms';
import { debounceTime } from 'rxjs/operators';

const formControl = new FormControl();

formControl.valueChanges.pipe(
    debounceTime(300) // Wait 300ms after the user stops typing
).subscribe(searchTerm => {
    console.log('Fetching data for:', searchTerm);
});
```

[`debounceTime` marbles](https://rxjsmarbles.dev/debounceTime)

## `catchError`

**Use Case**: Handling HTTP errors without breaking the stream.

```ts
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { EMPTY } from 'rxjs';

http = inject(HttpClient);

fetchData() {
  return this.http.get('api/data').pipe(
    catchError(error => {
         console.error('Error fetching data', error);
         return EMPTY; // Return an empty observable if error occurs
    })
  );
}
```

[`catchError` marbles](https://rxjsmarbles.dev/catchError)

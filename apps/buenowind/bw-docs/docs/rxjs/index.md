# RxJS

RxJS is a powerful library that gives us ready-to-use `Observables` and tons of helpful tools, called
`operators`, that make working with streams of data easier.

[RxJS Docs](https://rxjs.dev/)

If we already understand how Observables work at a basic level, we're well on our way to understanding
RxJS,because RxJS is really just Observables plus a bunch of extra superpowers.

One of the things RxJS gives us is something called operators.
These are functions that either create new Observables or take an existing Observable and change it
somehow.

[RxJS Marbles](https://rxjsmarbles.dev/combineLatest)

When we use an operator that changes an Observable, we do it by using the `.pipe()` method.
Inside `.pipe()`, we can chain operators like `map()` and `filter()` to transform the data coming
through the stream.

For example, if we have a stream of numbers from 1 to 5, and we use `map(value => value * 2)`,
each number gets doubled, turning into `2, 4, 6, 8, and 10`.
Then if we use `filter(value => value < 7)`, we only keep the values that are less than 7
 So the final result we see is `2, 4, and 6`.

Even though it might look like we're changing the original stream, we're not. The `pipe()` method creates
a new Observable based on the original one, with the transformations applied.
This helps us safely build new streams without affecting others, which is really useful when our app
starts getting more complex.

RxJS might seem a bit intimidating at first, but once we get used to it, it becomes an incredibly
powerful way to manage and react to data in our apps.

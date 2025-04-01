# `effect`

An `effect` in Angular works similarly to `computed` signals, but instead of calculating a value,
it runs arbitrary code whenever the signals it depends on change.

For example, we can create an effect to log changes to `comfortable` and `name`:

```ts
constructor() {
    effect(() => {
        console.log(this.gameState().comfortable());
        console.log(this.gameState().name());
    });
}
```

This `effect` will run once initially and then re-run whenever either `comfortable` or `name` changes.
If we only want to log one of the values, we can use separate effects:

```ts
constructor() {
    effect(() => {
        console.log(this.gameState().comfortable());
    });

    effect(() => {
        console.log(this.gameState().name());
    });
}
```

We can use effects to trigger actions like saving data or navigating when certain signals change.
For example, when the score in `gameState` changes, we might trigger an effect to save the score to
local storage.

Effects only run when a signal's value actually changes.
If a signal is set to its current value, the effect won't run.

`effect()` only runs when a signal's value changes, helping optimize performance by avoiding
unnecessary operations.

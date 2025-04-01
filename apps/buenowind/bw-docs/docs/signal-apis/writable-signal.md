# `signal`

## Creating a signal

```ts
gameState = signal({
    score: 0,
    level: 1,
});
```

This creates a signal `gameState` initialized with an object containing `score` and `level`.

## Accessing a signal

In the template, we can access the values using:

```html
{{ gameState().score }} - Level: {{ gameState().level }}
```

## Updating a signal's value

We can update the entire object in the signal using `set`:

```ts
this.gameState.set({ score: 100, level: 2 });
```

## Using `update` with an object

If we need to modify specific properties of the `gameState` object, we can use the `update` method:

```ts
increaseScore() {
    this.gameState.update(gameState => ({
        ...gameState,
        score: gameState.score + 10, // Increase score by 10
    }));
}

nextLevel() {
    this.gameState.update(gameState => ({
        ...gameState,
        level: gameState.level + 1, // Increase level by 1
    }));
}
```

The `update` method ensures that only the relevant property is changed, while the other properties
(like `score` or `level`) remain intact.

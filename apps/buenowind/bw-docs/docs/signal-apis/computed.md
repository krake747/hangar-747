# `computed`

In Angular, computed signals automatically update when the values they depend on change. For example:

```typescript
highScore = computed(() => this.gameState().score > 50);
```

This `highScore` signal updates whenever the score changes. We can also create a `levelMessage` signal:

```typescript
levelMessage = computed(() => {
    const level = this.gameState().level;
    return level === 1
        ? 'You are just starting!'
        : `Level ${level} is where the fun begins!`;
});
```

Lastly, a `gameMessage` signal combines both `highScore` and `levelMessage`:

```typescript
gameMessage = computed(() => {
    const highScoreText = this.highScore() ? 'You have a high score!' : '';
    return `${this.levelMessage()} ${highScoreText}`;
});
```

Computed signals automatically update when their dependent signals change.
For example, if the score or level changes, `gameMessage` will update accordingly.

Computed signals allow dynamic content that updates automatically when the signals they depend on change.
They cannot be manually updated and are recalculated based on dependencies,
making the UI more reactive and efficient.

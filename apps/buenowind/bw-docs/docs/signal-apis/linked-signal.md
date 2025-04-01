# `linkedSignal`

A `linkedSignal` is similar to a `computed` signal, but it has an important distinction:
a `linkedSignal` is **writable**, whereas a `computed` signal is **read-only**.
This means we can manually update the value of a `linkedSignal`, but it will still recompute based
on the signals it depends on.

Let’s say we have a game where we track the score and level.
We can use a `linkedSignal` to manage a `doubleScore` value, which is double the value of `score`.
We can manually update `doubleScore`, but it will always be recomputed if `score` changes.

```typescript
score = signal(10); // Player's score
doubleScore = linkedSignal({
    source: this.score,  // Source signal
    computation: (score) => score * 2  // Computation to double the score
});

increaseScore() {
    this.doubleScore.update((doubleScore) => doubleScore + 1); // Manually increase doubleScore
}

level = signal(1); // Player's level
levelMessage = computed(() => {
    return `You are at Level ${this.level()}`;
});
```

- **Writable**: `doubleScore` is a writable signal created using `linkedSignal`. This allows us to
  update its value manually.
- **Recomputation**: Even though we can manually adjust `doubleScore`, it will automatically recalculate
  whenever `score` changes (e.g., when we call `score.set(20)`).
- **Flexible**: The `linkedSignal` allows us to adjust the computed value, but it still respects the
- computation logic when the source signal (`score`) changes.

The `linkedSignal` is a powerful tool when we want both automatic updates and the flexibility to
modify a signal's value.

# `input`

In Angular, we can create reactive signals from input fields using the `input()` function.
We can consider a simple game scenario where the user inputs their current level and
experience points (XP). Based on this, we can calculate whether the player has leveled up.

We create input signals for `currentLevel` and `xpPoints`:

```typescript
currentLevel = input.required<number>();
xpPoints = input.required<number>();
```

We then create a computed signal to determine if the player has reached a new level based on their XP:

```ts
levelUp = computed(() => this.xpPoints() >= 1000);
```

Whenever the `xpPoints` changes, the `levelUp` signal will automatically update to reflect whether
the player has earned enough XP to level up.

We can use `effect()` to run some side effects whenever the player’s level or XP changes.
For example, we might want to log a message when the player levels up:

```ts
currentLevel = input.required<number>();
xpPoints = input.required<number>();

constructor() {
  effect(() => {
    if (this.levelUp()) {
      console.log(`Congratulations! You’ve leveled up to level ${this.currentLevel() + 1}!`);
    }
  });
}
```

This `effect` will trigger whenever `xpPoints` or `currentLevel` changes, and it will log a message
if the player reaches a new level.

Using `input()` in Angular allows us to easily bind user inputs to reactive signals.

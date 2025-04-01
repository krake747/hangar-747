# `model`

The `model()` function in Angular enables two-way data binding between a parent and child component,
allowing the child component to modify a signal's value while still receiving it from the parent.
Unlike `input()`, which is read-only in the child component, `model()` allows bidirectional changes
to the signal.

Let’s say we have a game component where the player's level and XP are part of the state:

```ts
import { Component, model } from '@angular/core';

@Component({
    selector: 'app-game',
    template: `<p>Your level: {{ currentLevel() }} and XP: {{ xpPoints() }}</p>`
})
export class GameComponent {
    currentLevel = model(1);
    xpPoints = model(0);
}
```

In this case, `currentLevel` and `xpPoints` are signals defined with `model()`, and their values can
be modified within the component or from the parent component. For example:

```ts
@Component({
    selector: 'app-parent',
    template: `<app-game [(currentLevel)]="playerLevel" [(xpPoints)]="playerXP"></app-game>`
})
export class ParentComponent {
    playerLevel = 1;
    playerXP = 0;
}
```

In this example, `currentLevel` and `xpPoints` are bound between the parent and child component using
two-way data binding. The parent component can update these signals, and any changes will automatically
propagate to the child component.

Using `model()`, we can both pass and modify the value of a signal between a parent and child component.
This is particularly useful for scenarios like custom form controls, where the signal needs to be
updated from both sides. However, for the majority of cases in Angular, it's recommended to focus
on one-way data flow, as it aligns better with declarative programming principles.

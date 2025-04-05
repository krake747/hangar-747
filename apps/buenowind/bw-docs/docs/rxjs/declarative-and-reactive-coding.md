# Reactive and declarative coding

In Angular, reactive and declarative coding focuses on how we handle data changes.
Instead of manually updating values (imperative), we define how values should change over time.

**Declarative**: We declare data and describe how it changes.
For example, we say "this is how I want my data to be, and it will update when necessary."

**Reactive**: We react to changes in data, like automatically fetching data when a user interacts
with the app.

!!! note

    **Transparent** Reactivity: Reacts to the latest value. We only see the current state, and updates
    happen automatically (e.g., Angular data binding).

    **Reified** Reactivity: Reacts to streams of changes over time. We work directly with observables
    and have explicit control over the flow of data (e.g., RxJS observables).

<!-- markdownlint-disable MD046 -->
```ts
tasks = ['Task 1', 'Task 2', 'Task 3'];
lastTask = tasks[tasks.length - 1];
```

This is declarative because we define what `lastTask` is just by looking at its declaration.

But when the data changes, we need a reactive way to update it:

```ts
tasks = signal(['Task 1', 'Task 2', 'Task 3']);
lastTask = computed(() => this.tasks()[this.tasks().length - 1]);
```

Now, if `tasks` changes, `lastTask` updates automatically.

```ts
tasks.update((tasks) => [...tasks, 'Task 4']);
```

With this reactive setup, `lastTask` will update whenever tasks changes.

Another example where we combine multiple streams (like user data and settings) to dynamically update
values:

```ts
name$ = combineLatest([this.userService.getName(), this.settingsService.getPreferences()])
.pipe(map(([name, preferences]) => preferences.fullName ? `${name.first} ${name.last}` : name.first));
```

With reactive code, we avoid manually handling updates, making our app more maintainable and easier
to reason about.

# Signals

Signals are a new feature in Angular that help  manage reactive data in a more efficient way.

Normally, Angular checks everything in our app to see if something has changed so it can update the view.
This is a bit slow because it checks a lot of things, even if nothing changed.
Signals flip this by **telling Angular when something has changed** instead of making Angular guess.
This helps make Angular faster and more efficient.

Here's how it works:

1. **Creating a signal**: We wrap a value with the `signal()` function, like `age = signal(25)`.
2. **Accessing a signal**: In the template, we access the signal's value with `age()`.
3. **Updating a signal**: We can update a signal's value using `age.set(26)`, which tells Angular
   that the value has changed.
4. **Computed signals**: We can create a signal that depends on others. For example,
   `isAdult = computed(() => this.age() >= 18)` automatically updates when the `age` signal changes.
5. **Effects**: Effects run whenever signals change and can be used to trigger actions, like
   updating the UI or logging a message. For example, we could use an effect to log when the `age` changes.

So, signals make it easier to manage reactive data and make Angular faster because they notify it
when something has changed, instead of Angular having to check everything all the time.

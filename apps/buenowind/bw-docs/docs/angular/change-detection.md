# Change detection

Signals in Angular make change detection simpler and more automatic.
In the old way, Angular needed to keep track of every possible change in the app to know when to update
things like the screen, which could be complicated and sometimes slow.
We had to manage this carefully, especially for performance.

But with signals, when something changes, **signals automatically notify Angular**
that it needs to update the screen. This makes things easier to handle and improves performance,
because signals make it more intuitive to track changes without manually managing them.

In the past, Angular used two main strategies for change detection:

1. **Default**: Angular checks everything all the time for any change.
2. **OnPush**: Angular only checks things when **specific actions** happen.

## `OnPush` vs `Default`

Angular needs to know when something in our app changes so it can update the UI
Instead of constantly checking everything, it uses a tool called `zone.js` to detect when certain
things happen that might cause a change, such as:

- A component is initialized
- Events being triggered (like a button click)
- Handling the response of an HTTP request
- MacroTasks such as `setTimeout()` and `setInterval()`
- MicroTasks such as handling Promises

When one of these events happens, Angular checks all components to see if they need updating.
This is how the `Default` change detection strategy works. Angular will check every component in the
component tree, even if only one actually changed.

`Default` strategy is generally fine in terms of performance.
Angular can perform these checks quite quickly.
Using `OnPush` change detection instead will still improve performance.

The `OnPush` change detection strategy in Angular is like telling Angular to only check for changes
in certain situations instead of checking everything all the time.

With `OnPush`, the component will only update if:

- Events being triggered (like a button click)
- The component's inputs change (but it must be a completely new value, not just a change to an
  existing object, the `reference` of the input value needs to change).
- The **async** pipe is used and the value it's watching changes.
  (The `async` uses under the hood `markForCheck` function)

If none of these conditions are met, Angular won't check the component, saving time and improving performance.
It's more efficient, but we need to be mindful of how we manage state and inputs to make sure
changes are detected.

It is also possible to manually inject the `ChangeDetectorRef` and trigger change detection yourself,
but this is generally best avoided and should only be treated as a last resort.

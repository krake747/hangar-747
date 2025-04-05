# Async pipe

In Angular, the **async pipe** is useful for working with observables in templates.
It automatically subscribes to streams and updates the view when new data comes in.
However, in apps using **signals**, we use `toSignal` instead of the async pipe to avoid some issues
with subscriptions.

Three key reasons to use the async pipe:

1. **Automatic Subscription Management**: It subscribes and unsubscribes from streams automatically.
2. **Keeps Data Reactive**: Using it ensures data stays in streams until it reaches the template.
3. **Triggers Change Detection**: When new data comes in, the async pipe triggers change detection,
   important for OnPush change detection.

## Unsubscribing from observables

When we work with observables, they can keep emitting values over time.
If we don't unsubscribe, it can cause memory leaks. Multiple subscriptions might keep running in the
background, slowing down the app.

Some observables emit once, others continuously. Either way, we need to unsubscribe when we're done
to avoid using extra memory. It's better to always unsubscribe, even if we think we don’t need to,
than to risk forgetting in a situation where we should have.

The `async pipe` helps by automatically unsubscribing when the component is destroyed.
If we subscribe manually, we need to remember to unsubscribe ourselves.

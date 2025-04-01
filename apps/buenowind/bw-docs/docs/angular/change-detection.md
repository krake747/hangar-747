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

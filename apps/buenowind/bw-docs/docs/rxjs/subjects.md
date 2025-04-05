# Subjects

A **Subject** is a special kind of Observable in RxJS that lets us both listen to data and push new
data into the stream.
This means that a `Subject` behaves like a regular Observable when we subscribe to it,
but it also lets us send new values into it by calling methods like `next()`.

To understand why `Subjects` are different, we first need to know about cold and hot observables.

A **cold observable** creates new data each time we subscribe.
In other words, every subscriber gets its own separate experience, like each person watching their
own copy of a video.

On the other hand, a **hot observable**, like a `Subject`, shares one live experience with all subscribers,
like tuning into a live broadcast. Everyone sees the same data at the same time.

This shared experience is what makes `Subjects` "hot" or "multicast". We have one source of data being
pushed out to many listeners.
With a regular Observable, the data usually comes from inside the Observable itself.
When we subscribe, the Observable starts running its logic to produce values.
But with a `Subject`, the data comes from **outside** the stream.
We create the `Subject` first and then push values into it from elsewhere in the application.
These values are immediately shared with all current subscribers.

For example, if we call `next(1)` on a `Subject` and only one subscriber is listening, only that
subscriber sees the value. If we add a second subscriber afterward, it won't see the earlier value
unless we emit another one.
This is different from an Observable, where subscribing always restarts the data flow.

One really useful reason to use a `Subject` is when we want to manually control the flow of data.
If we're working on a feature like a to-do list, we could use a `Subject` to hold our list of to-dos.
Every time we add a new to-do, we call `next()` to push the updated list into the stream,
and anyone listening gets the new list immediately.

## Behavior Subject

There's also a helpful variation called a `BehaviorSubject`, which keeps track of the last value
that was emitted. When a new subscriber joins, it immediately gets the current value, even if it
subscribed late. This is great for managing shared state, because everyone always starts with
the most recent data.

!!! note

    Finally, we need to remember one important difference:

    With a regular `Observable`, every subscriber runs the data-producing logic again, which can be
    inefficient if that logic is expensive, like making an HTTP request.

    With a `Subject`, we can run that logic once and share the result with everyone, which can save time
    and resources.

In short, we use `Subjects` when we want to control the stream of data from the outside,
share it across multiple subscribers, and sometimes maintain the latest value to give to new listeners.

# The observer pattern

An observable is just a version of the **observer pattern**, a classic design pattern.
We can think of it like this:
we have a thing that sends data (the producer), and we have one or more things that want
to receive that data (the observers). The observable connects the two.

We can imagine an observable as a function. That function accepts an **observer**, which is just an
object that knows what to do with new data, errors, or the end of the stream.
The observable then calls methods like `next`, `error`, and `complete` on the observer whenever
something happens.

So when we "subscribe", we're really just saying:

> Here's how we want to respond to this data—go ahead and start sending it.

Understanding observables as just a **connection between data producers and consumers**.

RxJS gives us tools to manage these observables in powerful ways, but the core idea is simple:

> One stream of data, multiple listeners, clear structure.

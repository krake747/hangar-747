# Optimizing for the wrong thing

At its core, the desire to add small abstractions often comes from overvaluing the
**DRY** (Don't Repeat Yourself) principle over isolation. This can lead to a focus on reducing the
number of lines of code (LOC) as a sign of good engineering, which then becomes a goal in itself.
However, a bit of repetition is a small price to pay for maintaining independence and flexibility in
the code.

There's no strict rule for when abstraction is the right choice, but the general heuristic is to
value isolation 3 to 5 times more than reducing repetition.

!!! tip

    We should consider abstracting something when we encounter it at least 3 times
    (or preferably more) in the codebase. This approach helps ensure the abstraction is useful for
    a wider range of use cases.

!!! warning

    However, we should avoid abstracting Angular or state management library code. Attempting to build
    our own version of Angular or a custom state management system can lead to significant problems
    in the future, making it better to stick with the established tools and patterns.

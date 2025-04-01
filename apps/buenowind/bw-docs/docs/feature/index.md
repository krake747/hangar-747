---
tags:
  - lazy
---

# Feature

With all the reusable types and building blocks in place, it's time to start
implementing actual use cases and business flows that will be used directly
by our users!

- Implementation in `feature/<feature-name>/` folder
- Complete isolation between features, features can't import from one another
- Always lazy loaded with the help of loadChildren() (pointing to feature route config) instead of
`loadComponent()` as that ensures uniform API and universal extendability
- **Black box** - can contain any kind of implementation and building blocks. Even if it gets dirty,
the isolation prevents the spread, and it won't affect the rest of the application
- Throw-away nature, because of isolation, it's easy to throw it away and
start over. On the more positive note, it also becomes easy to extract it
into a library or move around.
- Large features can be further divided into lazy sub features to make them
more manageable and improve bundling (2nd, 3rd, ... level navigation)
- Sharing logic between features (and sub-features) follows *“extract one
level up rule”* (e.g. into parent lazy feature, pattern, core or ui)

!!! tip

    We should always strive to keep our lazy features as clean as possible, but the isolation guarantee
    allows us to optimize for speed of delivery where we can afford to be a bit more pragmatic and
    focus on delivering value to the users instead of perfect implementation.

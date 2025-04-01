---
tags:
  - eager
  - lazy
---

# Pattern

The last main architecture type that we need to discuss is the `pattern` type.

!!! quote

    In a nutshell, the pattern is a type which consists of a pre-packaged combination of standalones
    and injectables which implement a specific reusable use case which is consumed in a lazy feature
    (or rarely > in layout) with a help of "drop in" component instead of a route!

- Implementation in `pattern/<pattern-name>/` folder
- Great for implementing of **cross-cutting business features** like **document manager**,
**approval process, change history (audit log), notes or comments** which could be dropped in many
lazy features
- One way dependency between patterns and features, only **features can consume patternsbut not the**
**other way around**
- Consumed through main **"drop in" component used in the templates of lazy features** (or rarely, layouts)
- Sharing logic between patterns follows the familiar *"extract one level uprule"* (into core or ui)

!!! info "Why call it a pattern?"

    It is to differentiate between the level of individual generic reusable UI components
    (e.g. button or popover) and the level of pre-packaged combination of such components, e.g. "context
    menu" (a multiple buttons in a popover)

Even though patterns aren't used with the help of Angular routing, we can still lazy load them with
ease with the help of the `@defer` blocks, but it is always important to consider the tradeoffs and
only opt in for such an approach if patterns is heavy (large bundle size) or is only used after some
specific user interaction performed in the consumer lazy feature.

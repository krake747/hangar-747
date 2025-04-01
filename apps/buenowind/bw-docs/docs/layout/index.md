---
tags:
  - eager
---

# Layout

With our core ready, our application can bootstrap and even kickstart some
processes, but it won't display anything to the user just yet.

Typically, most applications will use one or more layouts as a *"frame"* around
the slot in which we're going to display specific lazy features based on user navigation or interaction.

- Implementation in `layout/` folder
- Eager, available from start, part of the initial bundles size
- Only template context based (standalones) building blocks like components, directives, and pipes
- Can and most likely will consume services and state from `core/`
- Can and will most likely consumer some standalones from `ui/` like `Avatar`, `Button` or `Menu`
- In case of single main layout which is used for the whole app it will be used
directly in the template of the `AppComponent`
- In the case of multiple layouts, it will be used as part of the route config
- Applications can have more than one layout, e.g. login, signup, authenticated
- Some application may need a custom layout per feature and in that case the `layout/` folder
may end up empty, the `AppComponent` will then end up with a template consisting only of
single `<router-outlet />`

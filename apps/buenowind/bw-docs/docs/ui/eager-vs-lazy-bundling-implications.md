# Eager vs lazy bundling implications

Standalone UI components can be used in different parts of an app without making the initial load slower.
Since they are "cherry-picked" only when needed, the **bundler** (which prepares files for the browser)
handles them efficiently.

Here's how it works with an `Avatar` component (which shows a user's profile picture):

## Eager Only

If it's *only* used in the main layout (e.g., to show the logged-in user's avatar),
it will be bundled with the *main app files*. In this case, we might move it from the `ui/` to the
`layout/`.

## Eager & Lazy

If it's used *both* in the main layout *and* in multiple lazy-loaded sections,
it stays in the *main bundle* because it's needed early.

## Multiple Lazy Features

If it's *only* used in multiple lazy-loaded sections,
the bundler creates a *separate file for it*. The first time a user visitsone of these sections,
the file is loaded, and it's *shared across all features* without needing to reload.

## Single Lazy Feature

If it's *only* used in one lazy feature, it will be
*bundled with that feature*. If we later need it elsewhere, we can move it to `ui/` so other
features can use it.

## Summary

The same rules apply to directives and pipes, making sure we only load what we need,
when we need it!

!!! tip

    We should always make sure that the location of the implementation of any give standalone is as
    precise as possible. That's why if the component is only used in the layout
    the implementation should also be `layout/` folder.

# How to handle verbosity

Instead of creating unnecessary shortcuts in our code, we should stick to using Angular
and our chosen state management library as they are. While this might feel like extra typing,
Angular and NgRx, NgXs keep improving, making things easier over time.

For example, Angular 16 introduced `takeUntilDestroyed`, which removes the need for extra code
to clean up observables. NgRx also keeps simplifying its API with features like action groups
and functional effects.

Since some level of extra code is unavoidable, we can use tools to help:

**Schematics**: Automatically generate components, services, and state features using best practices.

**Consistent Naming**: Makes search-and-replace in our IDE more effective.

By relying on Angular's built-in improvements and these tools, we keep our code clean and
future-proof without adding unnecessary complexity.

---
tags:
  - eager
  - lazy
---

# UI

Now it's time to talk about the UI type which is going to live in the `ui/`
folder. This type is special in that it's basically the only type that will in practice be consumed by
both the eager and lazy parts of our application.

- Implementation in `ui/` folder
- Only template context based (standalones) building blocks like
components, directives, and pipes
- Eager / lazy - each individual standalone will be imported and used
explicitly in each feature, pattern or layout that needs it, the bundler is
then able to determine the optimal bundle into which such a standalone
will be bundled
- Only generic reusable UI components, directives and pipes which
**communicate only through inputs and outputs**
- Never bound to a specific state through service or state management
library (as it **can't import from the core** anyway)

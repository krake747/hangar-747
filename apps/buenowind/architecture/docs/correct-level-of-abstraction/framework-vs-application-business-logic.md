# Framework vs application business logic

In general, our primary focus and effort should be dedicated to developing the application's
business logic—the parts that deliver real value to users, like user flows, pages, and forms.
This is because we already have a powerful framework like Angular to handle the underlying structure,
and state management libraries like NgRx or NgXs to manage the state.

This approach contrasts sharply with creating complex systems like `MyOrgAngular+` a collection of
base components and tiny wrappers that combine logic in unexpected ways, like mixing selectors and actions.

While the downsides of such an approach might be limited in smaller organizations with only a few projects,
the issues grow much larger as the number of projects increases.
Over time, these abstractions tend to diverge across codebases, leading to maintenance challenges.

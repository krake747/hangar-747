# Common neative consequences of adding abstractions

Abstractions can generally be categorized into **business logic-related** and **framework-related** abstractions.

**Business logic-related abstractions** can sometimes evolve into a **framework within a framework.**
When business logic becomes overly abstract and complex, it can start to feel like using a
custom-built framework rather than Angular itself.

**Framework-related abstractions** (such as custom base components, forms, and directives) can
create issues when Angular evolves in a direction that doesn't align with the assumptions made
in the custom implementation. This often leads to significant effort in undoing and reworking
the custom abstraction, as it can "pollute" the rest of the business logic.

Custom abstractions can also break **automatic migrations** provided by tools like `ng update`.
A good example were custom `MyOrgIf` directives when Angular 17 introduced the new
[control flow](https://angular.dev/guide/templates/control-flow) syntax.

Over time, these custom abstractions can introduce various unexpected bugs and maintenance challenges.

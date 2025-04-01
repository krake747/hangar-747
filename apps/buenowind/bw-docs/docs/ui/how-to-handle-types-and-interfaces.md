# How to handle types and interfaces

In most projects, we face a common issue when working with **standalone UI components** because of
the following constraints:

**UI Components Use Inputs & Outputs** – They receive data through **inputs** and notify changes
through **outputs** instead of fetching data themselves.

**Type Safety** – To ensure type safety, we need to define **interfaces and types** for these inputs
and outputs.

**Data Comes from Services or State Management** – The actual data (with its types) is handled by
**services** or **state management** in the `core/` or specific `features/`.

**No Direct Dependencies** – UI components **must not** import anything from the `core/` or any
lazy `feature/` to stay reusable.

## How do we ensure type safety without breaking isolation?

### Reworking UI Standalone into a Pattern

One possible approach to solving the issue is to turn the `avatar` component into a pattern.
Patterns are a more advanced architectural type that can import types like the `User` interface from
the `core/`, making it possible to work around the existing constraints.

However, this approach comes with certain drawbacks. By converting the `avatar` component into a pattern,
it would gain access to core services and logic, which **contradicts the principle of isolation**.

Additionally, the `avatar` component is not truly a pattern but rather a generic UI component, meaning
this solution would be misusing the architectural type system to bypass the established constraints.
While this method is technically feasible, it is not an ideal approach because it **compromises the**
**proper separation of concerns**.

### Define interfaces and types in the ui standalone itself

To solve this challenge, the best approach is to define the types and interfaces directly inside the
UI component (like the `Avatar` component). This means that the component will know exactly what data
it needs and won't depend on other parts of the codebase.

This has several benefits:

- The component becomes self-contained and only cares about the data it needs to function.
- The types will be specific to that component, like creating an `Avatar` type instead of using a
  broader `User` type.
- Even if other parts of the code change, like how a `User` is structured, the `Avatar` component
  can stay the same.

However, this means that if the `User` and `Avatar` types get more different over time, we might
need to set up a mapping between them.

Another option is to break down the data further, such as using just `imageUrl`, `name`, and `role`
as simple types (like `string` or `number`). This avoids needing to reference the whole `User` type
and makes things simpler. For example:

!!! example

    === "Before"

        ```html
        <my-org-ui-avatar [imageUrl]="user.profileImageUrl">
        ```

    === "After"

        ```html
        <my-org-ui-avatar [imageUrl]="user.resource.imageUrl">
        ```

    This keeps the `Avatar` component flexible and makes it easier to adjust as the `User` changes.

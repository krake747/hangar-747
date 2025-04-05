# When does managing state become hard

Managing state gets harder as our app grows. Initially, it's simple, like toggling a greeting,
but when state needs to be shared across the entire app, we have to use a shared service.

As we expand, we might want to save the state (like the greeting preference) in storage to persist
it across app reloads. This means we need to load the state on app start, adding complexity.

Even in this simple scenario, we face challenges, like manually loading state and ensuring it's
updated properly. For bigger apps, we must handle things like conflicting updates, caching data, and
making sure updates are predictable, especially when multiple developers are involved.

As apps grow, state management becomes more complex to handle it effectively.

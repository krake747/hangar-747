# Standalone core

Imagine `provideCore({ routes })` as a single, centralized place in our application where we put
everything that needs to be set up or initialized when the app starts.

It's like a to-do list for everything the app needs to function properly, such as loading
configurations, setting up services, and triggering some important processes.

## The core things we put in `provideCore()`

### Angular's basic providers

These are the core Angular services that make our app run. Think of them as essential tools that
Angular needs to handle things like routing, HTTP requests, and animations. These are provided by
Angular in functions like `provideRouter()`, `provideHttpClient()`, and `provideAnimations()` and
usually supoprt passing of additional configuration to parameterize or enhance provided functionality.

### State management providers

Apps often have **state**, meaning information that is shared across different parts of the app
(like whether a user is logged in, or their shopping cart). **State management** libraries like NgRx
help us manage that state consistently.

`provideStore()` sets up the state container for the app, while `provideEffects()`
lets us handle things that happen in the background, like fetching data or updating the user's session.

### Third-party libraries for infrastructure

Apps often use extra tools that handle tasks like:

- **Logging**: Recording what happens inside the app, useful for debugging.
- **Translations**: Switching between different languages in the app.
- **Analytics**: Tracking how users interact with the app.

These providers help integrate those external tools.

### Initialization logic (Kick-starting processes)

Some things in our app need to be done before the app can fully function.
For example, we might need to fetch the user's authentication token or load their profile before
anything else happens.

We set this up in `provideCore()` using `provideEnvironmentInitializer()`, which allows us to
run code that needs to execute during startup. We could dispatch an event to load the user's data or
initialize some services.

### App-specific logic

This is where we set up things specific to our app, like business logic in the form of services, guards,
state management specific logic and more. It is something we need from the start.

For example, like managing authentication where we check if the user is logged in, storing their token,
and making sure they have access to certain features (like a protected page) or perform other backend
requests.

Or, we make sure the app's business logic is in place before users start interacting with it.

### App-specific core logic shared by multiple lazy features

Sometimes, certain logic needs to be used by more than one part of the app.
For example, imagine we have an `order feature` that handles all the order-related tasks.
Now, we want a `dashboard feature` to show some stats about those orders.

The problem is that since the order feature and the dashboard feature are **separated**,
the dashboard can't just use the order logic directly.

We move the order management logic to the **core**. The core acts as a shared space where both the
order and dashboard features can access the same services or data. This way, both features can work
with the same order data without directly depending on each other, keeping the features independent
but still able to share important functionality.

## Why is this important?

### Centralized setup

By having everything in `provideCore()` we don't have to worry about missing anything when
initializing the app. Everything that's crucial to the app's startup is in one place.

### Clean and manageable

We only need to call `provideCore()` in the `app.config.ts` file, making the setup very clean and
easy to manage. We avoid putting setup logic all over the place and ensure that things are initialized
in the right order.

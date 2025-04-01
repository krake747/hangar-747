# Grouping logic by domain

In an Angular app, it's important to organize the code in a way that makes it easy to maintain and scale.
One of the best practices is to **group logic by domain**, not by the type of building block
(like services or reducers).

For example:

- If we have logic related to **authentication** or **user management**, it should go into
`core/auth/` or `core/users/`.
- If we have **state management** for orders (like managing CRUD operations), it should be in `core/orders/`.

This approach makes it clearer which parts of the app relate to which business logic. It's easier to
understand the structure, especially when features grow and evolve.

## Why Domain-Based Grouping is Better

### Clarity

Organizing by domain makes it clear what each folder represents (e.g., `core/orders/` for all things
related to orders).

### Scalability

As our app grows, this structure lets us add new logic or features in a consistent and logical way,
without having to deal with a mix of services, reducers, and other building blocks in one folder.

## Specialized Folders

While domain-based grouping is generally preferred, there may be exceptions.
The **core** is also the place to store simple, reusable **utils** functions that don't belong
to any specific feature but are helpful across the app.

For example:

- **Interceptors** (like HTTP interceptors) might be grouped in `core/interceptors/`.
- **Generic utilities** that don't belong to any specific domain, such as date formatting or query
parameter parsing, could go in `core/utils/`.

## Summary

When organizing the code, **group by domain** to keep things clear, but be open to exceptions for
more specialized utilities like interceptors or date parsing, which might deserve their own folder.

# Generic reuseable UI components

A **generic UI component** is like the building blocks we find in libraries such as
Angular Material or PrimeNG. Things like `buttons`, `cards`, or `menus`. Most apps use some of these
libraries, but sometimes we need to create our own reusable UI components.

The key rule for generic UI components is that they **do not fetch data** directly from services
or state management tools. Instead, they **receive data as inputs** and **send updates as outputs**.
This makes them **fully reusable** anywhere in the app without creating unnecessary dependencies or complexity.

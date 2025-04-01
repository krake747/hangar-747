# Decorators

**Decorators** in Angular help attach extra information to classes (like components, directives, pipes,
and services) to define how they should behave or be used in the app

## `@Component`

A `@Component` decorator is used to define a component in Angular. It tells Angular how to create and
display the component, including its name (selector), template, and styles. It's like a special note
attached to the component class, telling Angular what this class is supposed to do.

## `@Directive`

The `@Directive` decorator is for creating custom directives. Unlike components, directives don't have
templates of their own. Instead, they modify the behavior of existing elements.

## `@Pipe`

The `@Pipe` decorator is used to create custom pipes in Angular. Pipes take input data, transform it
(like reversing a string), and then output it in a different format.

## `@Injectable`

The `@Injectable` decorator is used to define services that can be injected into components.
Services help separate business logic from UI components and can be shared across multiple components.
With `@Injectable`, Angular manages the service and its data, which makes it easy to reuse across
the app.

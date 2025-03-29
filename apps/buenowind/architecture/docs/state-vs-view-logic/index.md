# State vs view logic

It's a good idea to separate the logic in our app into two main parts:

**State handling or business logic**: This is where we manage the data or rules of the app, usually
handled by services or state management
[RxJs and Signals](https://modernangular.com/articles/state-management-with-rxjs-and-signals),
[NgRx](https://ngrx.io/) or [NgXs](https://www.ngxs.io/).

**View logic**: This is how we show the data in the user interface (UI), which is the job of the component.

Most components should be simple "view-only" components. They should just display data (like from a
service or NgRx selector) and not have much logic themselves. Any actions from the user
(like button clicks)should be handled by a service or state manager, not the component.

However, components that deal with things like forms or dialogs (such as an editor or a confirm dialog)
are exceptions. These types of components can have their own specific logic,
since they need to manage user interactions directly.

!!! warning

    The introduction of **Signals** provides a straightforward way to manage state in Angular using
    RxJS and signals, without the need for any third-party libraries. However, if the complexity
    grows and it becomes difficult to manage, it's advisable to consider using state management
    libraries like **NgRx** or **NgXs**. These libraries offer more robust solutions for
    handling state in larger applications.

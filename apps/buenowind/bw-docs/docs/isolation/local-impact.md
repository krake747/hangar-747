# Local impact

When we change a lazy-loaded feature in an app, it should only affect that feature and not other
parts of the app. This makes it easier for us as developers to understand and work with the code,
especially for new team members who can quickly start contributing.

However, if we make changes to important parts of the app like the core logic or shared UI elements,
those changes can affect many parts of the app. So, it's important for us to have good tests in place
to make sure everything still works as expected.

Isolation guaranteess local impact and local testing.

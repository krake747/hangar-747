# One way binding

## Property Binding `[]`

We use square brackets to bind the value of a class property (like `firstName`) to a property of an
HTML element (such as the `value` of an input). This binding ensures that the element’s property
is automatically updated whenever the class value changes.

For example, `<input [value]="firstName" />` sets the input field’s value to the `firstName` property.

## Event Binding `()`

We use parentheses to bind an event (like `click`) to a method in our class. When the event is triggered,
the associated method is executed. We can also pass the event details (using `$event`) to the method.

For example, `<button (click)="handleClick($event)">click me</button>` binds the `click` event to the
`handleClick` method.

## Interpolation `{{}}`

We use interpolation to display the result of an expression or the value of a class property directly
in the template. While interpolation is useful for displaying values, we should avoid using it for
property binding since it treats the value as a string.

For example, `<p>Hi, {{ name }}</p>` displays the value of `name` inside the paragraph.

# Embedded templates

An embedded template in Angular lets us define a chunk of HTML that Angular won’t immediately display.
It's like creating a "template stamp" that we can use dynamically when needed.
This is useful for things like showing loading states or conditionally rendering content.

Here’s an example of an embedded template:

```html
<ng-template #myTemplate>
  <p>Hello there</p>
</ng-template>
```

At first, nothing will display because Angular doesn’t render the template right away.
To use this template, we can reference it with `*ngTemplateOutlet`:

```html
<ng-container *ngTemplateOutlet="myTemplate"></ng-container>
```

This lets us "stamp out" the template in the DOM whenever we need it.

## Why Use Embedded Templates?

**No Extra DOM Elements**: Unlike a regular `<div>`, `<ng-container>` doesn’t add extra elements to
the DOM.

**Dynamic Rendering**: We can render the template multiple times or pass different data
(like changing a greeting message) without re-writing the template.

Example of passing dynamic data:

```html
<ng-template #myTemplate let-greeting="greeting">
  <p>{{ greeting }} there</p>
</ng-template>

<ng-container *ngTemplateOutlet="myTemplate; context: { greeting: 'Hi' }"></ng-container>
<ng-container *ngTemplateOutlet="myTemplate; context: { greeting: 'Hello' }"></ng-container>
<ng-container *ngTemplateOutlet="myTemplate; context: { greeting: 'Yo' }"></ng-container>
```

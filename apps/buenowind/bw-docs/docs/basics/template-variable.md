# Template variable

A template variable allows us to reference an element in the template.

For example, `<p #myParagraph></p>` creates a variable `myParagraph` that we can use to interact with
the paragraph element.

Here’s an example where we use the variable to change the content of the paragraph when a button is clicked:

```html
<button (click)="myParagraph.innerHTML = 'Hello there'">Click me</button>
```

Template variables are useful in more complex situations where we need to interact with or manipulate
elements in the template.

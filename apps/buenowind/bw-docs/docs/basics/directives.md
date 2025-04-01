# Directives

A **directive** in Angular is similar to a component, but it doesn't have its own template.
Instead, it **modifies the behavior** of existing elements or components in the template.

To create a directive, we use the `@Directive` decorator, and we define a `selector` to specify
which elements the directive will apply to.

For example, a directive with the selector `[bgRed]` will apply to any element that has the `bgRed` attribute.

Here's a simple directive that changes the background color of an element:

```ts
@Directive({
  selector: '[randomColor]',
  host: {
    '[style.backgroundColor]': 'color',
  },
})
export class RandomColor {
  color = `#${Math.floor(Math.random() * 16777215).toString(16)}`;
}
```

Then we use it like this:

```html
<new-component randomColor />
```

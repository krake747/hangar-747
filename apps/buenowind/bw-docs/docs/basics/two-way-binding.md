# Two way binding

Two-way data binding means changes in the class and the template will automatically sync with each other.
If we update the `firstName` in the class, the template (like an input field) will show the updated value,
and if we change the value in the input field, it will automatically update the `firstName` in the class.

To set this up, we can do it like this:

```html
<input [value]="user.login" (input)="user.login = $event.target.value">
```

Alternatively, Angular provides a shortcut using `ngModel`:

```typescript
<input name="login" [(ngModel)]="user.login" />
// This syntactic sugar for
<input name="login" [ngModel]="user.login" (ngModelChange)="user.login = $event.target.value" />
```

With `[(ngModel)]`, "banana-in-a-box", Angular handles both updating the value in the input field and
updating the class automatically.

As of Angular 17.2 we can use the `model()` signal instead.

```typescript
<input [(firstName)]="firstName" />
[...]
firstName = model.required<string>();
```

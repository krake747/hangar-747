# How to handle complex pipes

Imagine we have a pipe that does some complex work (like transforming data) and we want to use this
in both components and services.
But we **don't want to mix** how Angular handles templates (the view part) and logic (the service part).
This would make the app messy.

To solve this, we **split the logic** into two parts:

1. A **simple pipe** used in the template.
2. A **service** that holds the complex logic.

By doing this, the service logic can be used both in **components** (in templates) and **services**
(in backend-like logic). This separation makes things cleaner and easier to maintain.

Now, if we want to use this pipe in another component, **we can't directly import it**, because of
how we've set up our architecture.

To work around this, we **move the pipe logic out of the UI component** and handle it in the
**parent component**. The parent will do the transformation, and then just pass the final data
down to the UI component.

!!! example

    === "Example 1 - Options driven APIs"

        We use the pipe in the parent component to transform data before passing it to the child UI
        component (like a dropdown).

        ```html
        <my-org-ui-price [price]="product.price | myOrgPatternPricePipe" />
        ```

    === "Example 2 - Template driven APIs (better)"

        We use the pipe in the parent to transform each price before passing it down as an item
        inside the dropdown. The dropdown just gets the final, transformed data.

        ```html
        <my-org-ui-dropdown>
            @for (let price of prices) {
                <my-org-ui-dropdown-item [value]="price.id">
                    {{ price.value | myOrgPatternPricePipe }}
                </my-org-ui-dropdown-item>
            }
        </my-org-ui-dropdown>
        ```

So, instead of letting the UI component directly handle the pipe, we handle it in the parent,
keeping everything clean and following the rules of the architecture.

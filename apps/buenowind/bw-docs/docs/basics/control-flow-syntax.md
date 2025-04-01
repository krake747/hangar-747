# Control flow syntax

Angular 17 introduced control flow syntax. This feature allows us to use a new template syntax to
write control flow statements, like if/else, for, and switch, instead of using the built-in
structural directives (`*ngIf`, `*ngFor`, and `*ngSwitch`).

=== "@if"

    ```html
    @if (condition){
        <p>display this</p>
    } @else if (someOtherCondition){
        <p>display this instead</p>
    } @else {
        <p>last resort</p>
    }
    ```

=== "@for"

    ```html
    @for (item of items; track $index){
    <li>{{ item }}</li>
    }
    ```

=== "@switch"

    ```html
    @switch (condition) {
        @case (caseA) {
            <p>display this</p>
        }
        @case (caseB) {
            <p>display this instead</p>
        }
        @default {
            <p>last resort</p>
        }
    }
    ```

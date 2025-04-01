# Pattern vs UI, core and feature

A `pattern` is different from a `UI` component because it connects to a data source, such as a service
or state management selector, whereas a UI component can only receive data through inputs and outputs.

A `pattern` differs from `core` logic because it includes both UI and logic, while core logic consists
only of headless logic without any UI.

A `pattern` is also different from a `feature` because it is used as a drop-in component, such as
`<my-org-document-manager>`, whereas a feature is accessed through routing.

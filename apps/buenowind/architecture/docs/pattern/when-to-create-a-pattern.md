# When to create a pattern

A pattern should be created when multiple features share both behavior and UI in a way that makes
full reuse beneficial.

1. If the **behavior and UI are very different**, it is better to keep them separate to allow independent
   changes in the future, even if this means some duplication.
2. If the **behavior is similar but the UI is different**, the shared logic should go into the
   `core/` folder, while each feature designs its own UI.
3. If the **behavior is different but the UI is similar**, the shared UI should go into the
   `ui/` folder, while each feature handles its own behavior.
4. If the **behavior and UI are both similar**, a pattern should be created in the
  `pattern/` folder. This avoids duplication but requires making the pattern flexible enough to handle
   different feature needs.

In general, patterns are great for implementing of cross-cutting business features.

The characteristic feature of all these patterns is that they either work the same in every consumer
or receive a relatively small amount of configuration from the consumer lazy feature to work in a slightly
different way.

=== "Document Manager"

    Handles documents like user guides and invoices.

    It can support actions like download, edit, and delete.

=== "Approval Process"

    Manages approvals for things like orders or invoices.

    It can define rules such as required approvals and actions like approve or reject.

=== "Change History"

    Tracks and displays audit logs, such as order updates.

    It can control change types and actions like revert or view details.

=== "Notes/Comments"

    Manages notes for products or customers.

    It supports actions like add, edit, and delete, and can distinguish between public and private notes.

!!! important

    Patterns are ideal for reusable business features that work the same across consumers or require
    minimal configuration to adapt. They can be dropped into features easily.

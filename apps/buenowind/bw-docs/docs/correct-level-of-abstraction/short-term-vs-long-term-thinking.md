# Short term vs long term thinking

Introducing small abstractions might seem like a good idea in the short term, as it can save us a
bit of effort by abstracting away repetitive lines of code. However, what is often overlooked is
that these abstractions create coupling between parts of the codebase that could have remained
independent and isolated.

In the long term, maintaining isolation between unrelated parts of the code provides the flexibility
needed for the project to evolve. This is especially important in frontend development, where
requirements often change rapidly. Keeping parts of the codebase isolated ensures the project can
adapt and grow without unnecessary dependencies causing complications.

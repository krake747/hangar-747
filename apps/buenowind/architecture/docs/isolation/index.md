# Isolation

Usually, programmers try to avoid repeating code (the "Don't Repeat Yourself" or **DRY** principle).
But in frontend apps, **isolation** (keeping different parts of the app separate), is
**way more important**, about **3 to 5 times more valuable than DRY**, as it aims to reduce coupling.

Why? Because frontend requirements change all the time! If we try too hard to remove every bit
of repeated code by making shared "super-components" or "god-services," they quickly
become too complex and hard to maintain.

Instead, it's often **better to allow some small code duplication** inside different lazy-loaded features.
This keeps them **independent** so we can change one feature without accidentally breaking another.

Of course, we should still reuse generic things like buttons, badges, or common utilities.
But for **business logic**, we should wait until we see the same code **at least 3 times** before
deciding to abstract it. There's no one-size-fits-all rule, but in the frontend world,
keeping things separate is usually the better choice!

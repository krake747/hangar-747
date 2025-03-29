# Standalones only

The `ui/` folder should only contain **standalone components, directives, and pipes**.
Basically, things related to how the UI looks and behaves in templates.
It should **not** contain services or business logic.

Each standalone component **builds its own template context**, which means we only include the parts
we actually use. This keeps our JavaScript bundles small and efficient.

We can use these **UI components** in both lazy and eager-loaded parts of the app without problems.

If a UI component needs data from a service, the **parent component** (like a layout or feature component)
should **fetch the data and pass it down** via an input property.
This keeps UI components clean and focused on rendering.

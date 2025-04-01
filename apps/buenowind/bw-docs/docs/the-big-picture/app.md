---
tags:
  - eager
---

# App

The app type will consist of all files starting with `app.` prefix.

These files will mostly stay as they were generated with the help of Angular Schematics when first
creating the workspace with the help of `ng new`.

- `app.component.ts` in most cases will contain either the `<router-outlet />` or in case of multiple
layouts per route or `<my-org-layout />` component
- `app.config.ts` import and call `provideCore({ routes })` function from `core.ts`
- `app.routes.ts` top level routes config, will import and use `feature/<feature-name>/` routes configs
of the first level lazy features

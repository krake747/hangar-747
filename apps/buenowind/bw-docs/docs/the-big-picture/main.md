---
tags:
  - eager
---

# Main

The main type will consist only of the `main.ts` file which is used to bootstrap
the application. The file itself will stay as it was generated with the help of
Angular Schematics when first creating the workspace.

```ts
import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

bootstrapApplication(AppComponent, appConfig).catch((err) => console.error(err));
```

import { bootstrapApplication, enableDebugTools } from "@angular/platform-browser";
import { AppComponent } from "./app/app.component";
import { appConfig } from "./app/app.config";

bootstrapApplication(AppComponent, appConfig)
    .then((applicationRef) => {
        enableDebugTools(applicationRef.components[0]);
    })
    .catch((err) => console.error(err));

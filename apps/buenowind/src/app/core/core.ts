import { provideHttpClient, withInterceptors } from "@angular/common/http";
import { provideEnvironmentInitializer, provideZoneChangeDetection } from "@angular/core";
import { provideAnimationsAsync } from "@angular/platform-browser/animations/async";
import {
    Routes,
    provideRouter,
    withComponentInputBinding,
    withEnabledBlockingInitialNavigation,
    withInMemoryScrolling,
    withRouterConfig
} from "@angular/router";

export interface CoreOptions {
    routes: Routes;
}

export function provideCore({ routes }: CoreOptions) {
    return [
        provideZoneChangeDetection({ eventCoalescing: true }),
        provideAnimationsAsync(),
        provideHttpClient(withInterceptors([])),
        provideRouter(
            routes,
            withRouterConfig({ onSameUrlNavigation: "reload" }),
            withComponentInputBinding(),
            withEnabledBlockingInitialNavigation(),
            withInMemoryScrolling({
                anchorScrolling: "enabled",
                scrollPositionRestoration: "enabled"
            })
        ),

        // other 3rd party libraries providers

        // other application specific providers and setup

        // perform initialization, has to be last
        provideEnvironmentInitializer(() => {
            console.log("environment initialized");
        })
    ];
}

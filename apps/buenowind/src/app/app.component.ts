import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "bw-root",
    imports: [RouterOutlet],
    template: `
        <h1>Welcome to {{ title }}!</h1>
        <h1 class="text-3xl font-bold underline">Hello BuenoWind!</h1>
        <router-outlet />
    `,
    styles: []
})
export class AppComponent {
    title = "buenowind";
}

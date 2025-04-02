import { ChangeDetectionStrategy, Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { NiceComponent } from "./nice/nice.component";

@Component({
    selector: "bw-home",
    imports: [RouterOutlet, NiceComponent],
    template: `
        <h1>Welcome to {{ title }}!</h1>
        <h1 class="text-3xl font-bold underline">Hello BuenoWind!</h1>
        <bw-nice />
        <router-outlet />
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class HomeComponent {
    title = "home";
}

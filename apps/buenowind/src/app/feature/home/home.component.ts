import { ChangeDetectionStrategy, Component } from "@angular/core";

@Component({
    selector: "bw-home",
    imports: [],
    template: `
        <h1>Welcome to {{ title }}!</h1>
        <h1 class="text-3xl font-bold underline">Hello BuenoWind!</h1>
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

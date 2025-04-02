import { ChangeDetectionStrategy, Component, input } from "@angular/core";

@Component({
    selector: "bw-pretty-name",
    imports: [],
    template: `
        <p class="bg-gray-200 text-amber-700">Pretty {{ name() }}</p>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class PrettyNameComponent {
    name = input("");
}

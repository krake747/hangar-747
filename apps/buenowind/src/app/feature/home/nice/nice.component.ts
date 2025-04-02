import { ChangeDetectionStrategy, Component } from "@angular/core";

@Component({
    selector: "bw-nice",
    imports: [],
    template: `
        <p>nice works!</p>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class NiceComponent {}

import { ChangeDetectionStrategy, Component } from "@angular/core";

@Component({
    selector: "bw-john",
    imports: [],
    template: `
        <p>john works!</p>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class JohnComponent {}

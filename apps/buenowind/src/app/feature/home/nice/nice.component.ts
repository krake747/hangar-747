import { ChangeDetectionStrategy, Component } from "@angular/core";
import { ButtonDirective } from "../../../ui/buttons/button.directive";

@Component({
    selector: "bw-nice",
    imports: [ButtonDirective],
    template: `
        <p>nice works!</p>
        <button bwButton>Hello</button>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class NiceComponent {}

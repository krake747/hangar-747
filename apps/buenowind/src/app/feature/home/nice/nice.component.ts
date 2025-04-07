import { ChangeDetectionStrategy, Component } from "@angular/core";
import { ButtonDirective } from "../../../ui/buttons/button.directive";
import { FaIconDirective } from "../../../ui/fa-icon.directive";

@Component({
    selector: "bw-nice",
    imports: [ButtonDirective, FaIconDirective],
    template: `
        <p>nice works!</p>
        <button bwButton variant="none" class="">
            <i bwFaIcon type="solid" icon="mug-hot"></i>
            Hello 1
        </button>
        <button bwButton variant="none" class="hover:text-amber-800">
            <i bwFaIcon type="solid" icon="mug-hot" class="group size-4 shrink-0"></i>
            Hello 2
        </button>
        <button bwButton variant="none" class="bg-red-500 text-white" [disabled]="true">
            Disabled
        </button>
    `,
    styles: `
        :host {
            display: block;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class NiceComponent {}

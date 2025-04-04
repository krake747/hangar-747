import { computed, Directive, input } from "@angular/core";
import { cva, VariantProps } from "class-variance-authority";
import { twMerge } from "tailwind-merge";

export type FaIconVariants = VariantProps<typeof iconVariants>;

const iconVariants = cva("inline-block", {
    variants: {
        type: {
            solid: "fas",
            regular: "far",
            brands: "fab"
        }
    },
    defaultVariants: {
        type: "solid"
    }
});

@Directive({
    selector: "[bwFaIcon]",
    standalone: true,
    host: {
        "[class]": "computedClass()"
    }
})
export class FaIconDirective {
    icon = input<string>();
    type = input<FaIconVariants["type"]>();

    computedClass = computed(() =>
        twMerge(iconVariants({ type: this.type() }), `fa-${this.icon()}`)
    );
}

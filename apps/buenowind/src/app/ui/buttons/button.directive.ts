import { computed, Directive, input } from "@angular/core";
import { cva, VariantProps } from "class-variance-authority";
import { twMerge } from "tailwind-merge";

export type ButtonVariants = VariantProps<typeof butttonVariants>;

const butttonVariants = cva(
    [
        "flex cursor-pointer items-center justify-center gap-2 rounded-lg font-semibold whitespace-nowrap outline-teal-500 transition-colors",
        "focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-teal-500",
        "active:scale-[0.98]",
        "disabled:pointer-events-none disabled:bg-gray-100"
    ],
    {
        variants: {
            variant: {
                none: "",
                primary: [
                    "bg-teal-600 text-white shadow-sm",
                    "hover:bg-teal-700",
                    "disabled:border-gray-300 disabled:bg-gray-100 disabled:text-gray-400"
                ],
                secondary: [
                    "border border-teal-300 bg-white text-teal-700 shadow-sm",
                    "hover:border-teal-300 hover:bg-teal-50 hover:text-teal-800",
                    "disabled:border-gray-200 disabled:bg-gray-50 disabled:text-gray-400"
                ],
                "secondary-gray": [
                    "border border-gray-300 bg-white text-gray-700 shadow-sm",
                    "hover:border-gray-300 hover:bg-gray-50 hover:text-gray-800",
                    "disabled:border-gray-200 disabled:bg-white disabled:text-gray-400"
                ],
                tertiary: ["bg-transparent text-teal-700", "hover:bg-teal-50 hover:text-teal-800"]
            },
            size: {
                sm: ["text-sm/5", "px-3 py-2"],
                md: ["text-sm/5", "px-3.5 py-2.5"],
                lg: ["text-base/6", "px-4 py-2.5"],
                xl: ["text-base/6", "px-4.5 py-3"],
                "2xl": ["text-lg/7", "px-5.5 py-4"]
            }
        },
        defaultVariants: {
            variant: "primary",
            size: "sm"
        }
    }
);

/**
 * Directive to apply consistent button styles using Tailwind CSS and variants.
 *
 * Usage:
 * ```html
 * <button bwButton variant="primary" size="sm" class="w-full">Click me</button>
 * ```
 *
 * Features:
 * - Supports `variant` and `size` inputs to control button styling via `buttonVariants`.
 * - Allows additional Tailwind or custom classes through the `class` input.
 * - Automatically merges and resolves class conflicts using `twMerge`.
 *
 * This directive promotes reusability and visual consistency across all buttons.
 */
@Directive({
    selector: "[bwButton]",
    standalone: true,
    host: {
        "[class]": "computedClass()"
    }
})
export class ButtonDirective {
    /**
     * The size of the button.
     * Accepts predefined sizes from the ButtonVariants type, such as 'sm', 'md', or 'lg'.
     */
    size = input<ButtonVariants["size"]>();

    /**
     * The visual variant of the button.
     * Accepts values like 'primary', 'secondary', 'outline', etc., defined in ButtonVariants.
     */
    variant = input<ButtonVariants["variant"]>();

    /**
     * Additional CSS classes to apply to the button.
     * Useful for extending or overriding default styles with Tailwind or custom classes.
     */
    class = input<string>();

    /**
     * Computes the final CSS class string for the button.
     *
     * - Merges the base styles from `buttonVariants` based on the current `variant` and `size` inputs.
     * - Applies any additional custom classes provided via the `class` input.
     * - Uses `twMerge` to intelligently merge Tailwind CSS classes, avoiding conflicts.
     *
     * This ensures the button has consistent styling while allowing customization.
     */
    readonly computedClass = computed(() => {
        return twMerge(
            butttonVariants({ variant: this.variant(), size: this.size() }),
            this.class()
        );
    });
}

import styled from "styled-components";

import { COLORS } from "./constants";

type ButtonVariant = "fill" | "outline" | "ghost";

type ButtonSize = "small" | "medium" | "large";

type ButtonProps = { variant: ButtonVariant; size: ButtonSize; children: any };

const SIZE_STYLES: Record<ButtonSize, Record<string, string>> = {
    small: {
        "--height": "52px",
        "--padding": "0 24px",
        "--font-size": `${18 / 16}rem`,
        "--border-radius": "2px"
    },
    medium: {
        "--height": "65px",
        "--padding": "0 36px",
        "--font-size": `${21 / 16}rem`,
        "--border-radius": "4px"
    },
    large: {
        "--height": "80px",
        "--padding": "0 48px",
        "--font-size": `${24 / 16}rem`,
        "--border-radius": "6px"
    }
};

// Template Literal Syntax which allows us to define CSS styles inside a template literal
// It allows us to use regular CSS and JavaScript expressions
const ButtonBase = styled.button`
    display: block;
    height: var(--height);
    padding: var(--padding);
    font-size: var(--font-size);
    border-radius: var(--border-radius);
    border: none;
    background: transparent;
    text-transform: uppercase;
    outline-offset: 3px;
    box-sizing: border-box;
`;

// This is a way to extend or compose components in a declarative manner.

// styled(ButtonBase) is a higher-order function that takes a ButtonBase component and returns a new component with
// added styles. FillButton will inherit the ButtonBase styles and then override its styles.
// FillButton behaves like ButtonBase but with the added styles.
export const FillButton = styled(ButtonBase)`
    background: ${COLORS.primary};
    color: ${COLORS.white};
    outline-color: ${COLORS.primary};

    &:hover {
        background: ${COLORS.primaryLight};
    }
`;

const OutlineButton = styled(ButtonBase)`
    background: ${COLORS.white};
    border: 2px solid ${COLORS.primary};
    color: ${COLORS.primary};
    outline-color: ${COLORS.primary};

    &:hover {
        background: ${COLORS.offwhite};
    }
`;

const GhostButton = styled(ButtonBase)`
    color: ${COLORS.gray};
    outline-color: ${COLORS.gray};

    &:hover {
        background: ${COLORS.transparentGray15};
        color: ${COLORS.black};
    }
`;

const COMPONENT_VARIANTS: Record<ButtonVariant, React.ElementType> = {
    fill: FillButton,
    outline: OutlineButton,
    ghost: GhostButton
};

const Button = ({ variant, size, children }: ButtonProps) => {
    const Component = COMPONENT_VARIANTS[variant];
    const styles = SIZE_STYLES[size];

    return <Component style={styles}>{children}</Component>;
};

export default Button;

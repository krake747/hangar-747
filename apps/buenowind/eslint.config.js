// @ts-check
const eslint = require("@eslint/js");
const tseslint = require("typescript-eslint");
const angular = require("angular-eslint");
const boundaries = require("eslint-plugin-boundaries");
const eslintPluginPrettierRecommended = require("eslint-plugin-prettier/recommended");

module.exports = tseslint.config(
    {
        files: ["**/*.ts"],
        plugins: {
            boundaries
        },
        extends: [
            eslint.configs.recommended,
            ...tseslint.configs.recommended,
            ...tseslint.configs.stylistic,
            ...angular.configs.tsRecommended,
            boundaries.configs.strict,
            eslintPluginPrettierRecommended
        ],
        processor: angular.processInlineTemplates,
        rules: {
            "boundaries/element-types": [
                "error",
                {
                    default: "disallow",
                    rules: [
                        {
                            from: "main",
                            allow: ["app"]
                        },
                        {
                            from: "core",
                            allow: ["env", "core"]
                        },
                        {
                            from: "ui",
                            allow: ["env", "ui"]
                        },
                        {
                            from: "layout",
                            allow: ["env", "ui", "pattern", "layout"]
                        },
                        {
                            from: "app",
                            allow: ["env", "app", "core", "layout", "feature-routes"]
                        },
                        {
                            from: ["pattern"],
                            allow: ["env", "core", "ui", "pattern"]
                        },
                        {
                            from: ["feature"],
                            allow: ["env", "core", "ui", "pattern", "feature"]
                        },
                        {
                            from: ["feature-routes"],
                            allow: ["env", "core", "pattern", "feature", "feature-routes"]
                        }
                    ]
                }
            ],
            "@angular-eslint/directive-selector": [
                "error",
                {
                    type: "attribute",
                    prefix: "bw",
                    style: "camelCase"
                }
            ],
            "@angular-eslint/component-selector": [
                "error",
                {
                    type: "element",
                    prefix: "bw",
                    style: "kebab-case"
                }
            ]
        },
        settings: {
            "import/resolver": {
                typescript: {
                    alwaysTryTypes: true
                }
            },
            "boundaries/ignore": [],
            "boundaries/dependency-nodes": ["import", "dynamic-import"],
            "boundaries/elements": [
                {
                    type: "env",
                    mode: "file",
                    pattern: "src/environments/environment*.ts",
                    capture: ["environments"]
                },
                {
                    type: "main",
                    mode: "file",
                    pattern: "src/main.ts",
                    capture: ["main"]
                },
                {
                    type: "app",
                    mode: "file",
                    pattern: "src/app/app(-|.)*.ts",
                    capture: ["app"]
                },
                {
                    type: "core",
                    pattern: "src/app/core",
                    capture: ["core"]
                },
                {
                    type: "ui",
                    pattern: "src/app/ui",
                    capture: ["ui"]
                },
                {
                    type: "layout",
                    pattern: "src/app/layout",
                    capture: ["layout"]
                },
                {
                    type: "pattern",
                    pattern: "src/app/pattern",
                    capture: ["pattern"]
                },
                {
                    type: "feature-routes",
                    mode: "file",
                    pattern: "src/app/feature/*/*.routes.ts",
                    capture: ["feature"]
                },
                {
                    type: "feature",
                    pattern: "src/app/feature/*",
                    capture: ["feature"]
                }
            ]
        }
    },
    {
        files: ["**/*.html"],
        extends: [...angular.configs.templateRecommended, ...angular.configs.templateAccessibility],
        rules: {}
    }
);

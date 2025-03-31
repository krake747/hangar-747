// @ts-check
const eslint = require("@eslint/js");
const tseslint = require("typescript-eslint");
const angular = require("angular-eslint");
const boundaries = require("eslint-plugin-boundaries");

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
            boundaries.configs.strict
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
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }]
                            ]
                        },
                        {
                            from: "ui",
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["ui", { app: "${from.app}" }]
                            ]
                        },
                        {
                            from: "layout",
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }],
                                ["ui", { app: "${from.app}" }],
                                ["pattern", { app: "${from.app}" }],
                                ["layout", { app: "${from.app}" }]
                            ]
                        },
                        {
                            from: "app",
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["app", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }],
                                ["layout", { app: "${from.app}" }],
                                ["feature-routes", { app: "${from.app}" }]
                            ]
                        },
                        {
                            from: ["pattern"],
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }],
                                ["ui", { app: "${from.app}" }],
                                ["pattern", { app: "${from.app}" }]
                            ]
                        },
                        {
                            from: ["feature"],
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }],
                                ["ui", { app: "${from.app}" }],
                                ["pattern", { app: "${from.app}" }],
                                ["feature", { app: "${from.app}", feature: "${from.feature}" }]
                            ]
                        },
                        {
                            from: ["feature-routes"],
                            allow: [
                                ["env", { app: "${from.app}" }],
                                ["core", { app: "${from.app}" }],
                                ["pattern", { app: "${from.app}" }],
                                ["feature", { app: "${from.app}", feature: "${from.feature}" }],
                                ["feature-routes", { app: "${from.app}", feature: "!${from.feature}" }]
                            ]
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

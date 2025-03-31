# Buenowind

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.2.5.

## Setup

### Angular

Install Angular CLI

```shell
npm i -g @angular/cli
```

Create BuenoWind application.

```shell
ng new buenowind --prefix bw --routing true --style css  --inline-template --inline-style
```

Create environment files.

```shell
ng g environments
```

Additional configuration during the `ng new` prompt:

We use plain CSS instead of SCSS to avoid extra build steps and because TailwindCSS provides the necessary styling capabilities. Additionally, we skip SSR and SSG support since our application requires user authentication, making client-side rendering the preferred approach.

To update all packages after scaffolding the project, we run the following commands sequentially:

```shell
ng update
npx npm-check-updates -u
```

We update the tsconfig with some extra configuration

```json
{
    "compilerOptions": {
        "noImplicitAny": true,
        "erasableSyntaxOnly": true,
    }
}
```

`noImplicitAny` setting forces TypeScript to require explicit type annotations for variables and function parameters where the type cannot be inferred.

`erasableSyntaxOnly` marks enums, namespaces and class parameter properties as errors. These pieces of syntax are not considered erasable. [erasableSyntaxOnly](https://www.totaltypescript.com/erasable-syntax-only)

### TailwindCSS

We have a dependency on [TailwindCSS](https://tailwindcss.com/) for styling our application.

We follow the [Angular Framework guide](https://tailwindcss.com/docs/installation/framework-guides/angular) to install it.

```shell
npm i tailwindcss @tailwindcss/postcss postcss --force
```

In addition, we include two other dependencies `class-variance-authority` and `tailwind-merge`.
We use them to efficiently manage and compose Tailwind CSS classes in our UI components to ensure cleaner,
more maintainable styling across the project.

- `class-variance-authority` helps define consistent component variants.
- `tailwind-merge` intelligently merges and deduplicates classes to prevent conflicts.

```shell
npm i class-variance-authority
npm i tailwind-merge
```

We also added basic CSS reset code in the `styles.css`.
It is from [Josh Comeau's A Modern CSS Reset](https://www.joshwcomeau.com/css/custom-css-reset/).

## Tools

As an Angular project grows, increasing coupling and technical debt can slow development and introduce unexpected bugs. Maintaining a scalable architecture helps preserve development speed and stability.

Building a strong foundation together with automated, tooling based architecture validation can preserve development
velocity in the long run.

We want to avoid "Hope-based" architecture. If we ignore the support of dedicated automated tooling, we're just hoping
that it stays that way and everyone involved will pay attention and review perfectly 100% of the time, which is just highly unlikely.

### VS Code

We add the `buenowind.code-workspace` file to centralize common configuration settings and recommended VS Code extensions.
This helps streamline setup, ensuring we all have the same tools for a consistent development environment.

### Prettier

We add a `.prettierrc` file to enforce consistent code formatting across the project.
This helps maintain clean, readable code and reduces formatting-related discrepancies.
This ensures a smoother development workflow with fewer stylistic conflicts

```shell
npm i -D prettier
```

We also add the offical VS Code extension to endhance our editor setup with Intellisence.
Then we add a prettier plugin to have the offical TailwindCSS class sorting.

We install it via `npm`.

```shell
npm i -D prettier prettier-plugin-tailwindcss
```

Then we add the plugin to the `.prettierrc` file.

```json
{
    "tailwindFunctions": ["cva", "tw"],
    "plugins": ["prettier-plugin-tailwindcss"]
}
```

Note: `"tw"` is optional and if we use it then it is needed for this utility helper function below.

```ts
export const tw = (strings: TemplateStringsArray, ...values: TemplateStringsArray[]) =>
    String.raw({ raw: strings }, ...values);
```

We also might want to add a `.prettierignore` file.

### Docker

We add support for Docker with `compose`, `Dockerfile`, `.dockerignore` and `nginx.conf` files.

```shell
docker compose up buenowind
```

We need to add two npm commands to the `pacage.json`

```json
{
    "build:ci": "ng build --configuration production",
    "test:ci": "ng test --watch=false --browsers=ChromeHeadless",
}
```

### ESlint

We can the following command to analyze our code for potential issues and enforce coding standards.

```shell
ng lint # This installs ESLint and the Angular plugin
```

[Setting Up ESLint and Prettier in Angular with VS Code and WebStorm](https://senoritadeveloper.medium.com/setting-up-eslint-and-prettier-in-angular-with-vs-code-and-webstorm-4be8d558caea)

This will install three npm packages.

- `angular-eslint`
- `eslint`
- `typescript-eslint`

### Automated Architecture Validation

This requires `ESLint`.

```shell
npm i -D eslint-plugin-boundaries eslint-import-resolver-typescript
```

We then do the setup in `eslint.config.js`. (git commit "tools(apps): add automated architecture validation")

### Bundle size analysis (and budgets)

**Bundle size matters** because a large bundle slows down how fast our app loads for users and increases the time it takes developers to build and test changes.

Tools like `@angular-experts/hawkeye` help us see what's inside our bundle. See [Hawkeye](https://angularexperts.io/blog/hawkeye-esbuild-analyzer)

```shell
npx @angular-experts/hawkeye init
```

This will update our `package.json` with the new `analyze` command (we might need to fix a typo to `-y`)

### Dependency graph analysis

The idea is to keep the code organized by ensuring that dependencies between different parts of the system are clean and straightforward.
This makes the code easier to maintain and extend.
We use automated tools to check that the dependencies stay clean at a high level,
but sometimes problems like circular dependencies can still sneak in.
We want to achieve a general sense of "one-way-ness" clean dependency graph.

`madge` is one of the best tools that we can use to evaluate the overall health of the code base from the perspective of its dependency graph.
With this tool we can generate dependency graphs for analyzing our project structure.

Additional dependency, [`Graphviz`](https://www.graphviz.org/) is required in order to generate
visual graphs which is the best way to read `madge` output.

```shell
sudo apt-get install graphviz
npm install -D madge npm-run-all
```

We update our `package.json` with some scripts. They start with `analyze:deps`.

### Angular Schematics

We update the schematics to include in the `angular.json`.

```json
{
    "@schematics/angular:component": {
        "inlineTemplate": true,
        "inlineStyle": true,
        "changeDetection": "OnPush",
        "style": "css"
    }
}
```

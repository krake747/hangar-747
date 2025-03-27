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

Additional configuration:

- CSS
- No SSR and SSG support

To update all packages we run these two commands in sequentially

```shell
npx ng update
npx npm-check-updates -u
```

We update the tsconfig with

```json
{
    "compilerOptions": {
        "noImplicitAny": true,
        "erasableSyntaxOnly": true,
    }
}
```

### TailwindCSS

We have a dependency on [TailwindCSS](https://tailwindcss.com/) for styling our application.

We follow the [Angular Framework guide](https://tailwindcss.com/docs/installation/framework-guides/angular).

```shell
npm i tailwindcss @tailwindcss/postcss postcss --force
```

In addition, we include two other dependencies `class-variance-authority` and `tailwind-merge` to better manage how we
create UI components with TailwindCSS.

TODO: Why we need these two?

```shell
npm i class-variance-authority
npm i tailwind-merge
```

We need the offical VS Code extension to endhance our editor setup with Intellisence and a prettier plugin
to have class sorting. We update our prettier with

```json
{
    "tailwindFunctions": ["cva", "tw"],
    "plugins": ["prettier-plugin-tailwindcss"]
}
```

Note: `"tw"` is needed for this utility helper function below.

```ts
export const tw = (strings: TemplateStringsArray, ...values: TemplateStringsArray[]) =>
    String.raw({ raw: strings }, ...values);
```

We also added basic CSS reset code in the `styles.css`.
It is from Josh Comeau [A Modern CSS Reset](https://www.joshwcomeau.com/css/custom-css-reset/).

## Tools

### VS Code

We add a `buenowind.code-workspace` file for some common configuration settings and VS code extensions.

### Prettier

We add a `.prettierrc` file for standardizing our formatting, a prettier dev dependency and some `npm` scripts.

```shell
npm i -D prettier
```

### Docker

We add support for Docker with `compose`, `Dockerfile`, `.dockerignore` and `nginx.conf` files.

```shell
docker compose up buenowind
```

### ESlint

We can the following command to analyze our code for potential issues and enforce coding standards.

```shell
ng lint #TODO
```

### Build Visualizers

We can install tools to generate visual representations of our bundle size and dependencies for better optimization
insights.

```shell
npm i -D esbuild-visualizer source-map-explorer http-server #TODO
```

### Dependency graph analysis

We can generate dependency graphs for analyzing our project structure.

```shell
sudo apt-get install graphviz # TODO
npm install -D madge npm-run-all
```

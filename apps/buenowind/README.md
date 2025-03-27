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
ng new buenowind --prefix bw --routing=true --inline-template --inline-style
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

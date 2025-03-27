# Buenowind

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.2.5.

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

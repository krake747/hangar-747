# TailBlaze

This sample app is used to demonstrate the installation process of Tailwind CSS v4 in a Blazor WASM app.

```shell
dotnet sln Hangar747.sln add TailBlaze/TailBlaze.csproj
```

## Installing Tailwind

We strip the Blazor WASM starter template from Bootstrap and remove the generated static files and the `.razor` files.

### NPM

This requires `npm`.

We follow the Tailwind CSS [Tailwind](https://tailwindcss.com/docs/installation/tailwind-cli) installation guide

```shell
npm install tailwindcss @tailwindcss/cli
npm install -D prettier prettier-plugin-tailwindcss
```

and the import to the main `app.css` file

```css
@import "tailwindcss";
```

We then add two scripts to the `package.json` file

- `build:css` allows us to build and minify the Tailwind CSS stylesheet
- `watch:css` allows us to `watch` the Tailwind CSS alongside `dotnet watch` (hot-reload)

```shell
"scripts": {
    "build:css": "npx @tailwindcss/cli -i ./wwwroot/css/app.css -o wwwroot/styles.min.css --minify",
    "watch:css": "npx @tailwindcss/cli -i ./wwwroot/css/app.css -o wwwroot/styles.min.css --minify --watch"
}
```

In order to build the style file everytime we add to the `.csproj` file.

```cs
<Target Name="Tailwind" BeforeTargets="PreBuildEvent">
    <Exec Command="npm run build:css" />
</Target>
```

To use the `app.min.css` file we need to add a link to the `index.html` file

```html
<link href="styles.min.css" rel="stylesheet"/>
```

## Local Dev

Execute from root directory the following command to run the app

```shell
dotnet run --project TailBlaze
```

### Watch Mode

We navigate to the TailBlaze directiory and in one terminal we run `dotnet watch` and in the other `npm run watch:css`.

# Eager Part

The **eager part** of the application refers to the code that loads initially. In practical terms,
this corresponds to the eager `main.js` bundle (or, with tools like `esbuild`,
the additional `chunk.js` files).

Any modules imported using the standard `import`/`export` JavaScript syntax are included
in these eager bundles.

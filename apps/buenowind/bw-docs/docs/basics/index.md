# Basics

## Browser

When a user requests the application, the browser accesses the specified URL. The server then
responds by returning the `index.html` file, which references an eager JavaScript bundle.

The browser downloads this eager bundle, with network conditions influencing the loading time.
To improve efficiency, bundles should be **gzip-compressed**, and file hashes should be used
to enable safe caching.

Once downloaded, the JavaScript is parsed and executed. This step is often the primary performance
bottleneck, particularly if the eager bundle is too large.

Additional feature bundles are loaded **dynamically** based on the URL or user interaction.
Like the initial bundle, these are also affected by network speed, parsing, and execution times.
Keeping these bundles small enhances overall application performance.

## Bundler

Angular projects consist of many TypeScript files that reference each other using imports and exports.
While modern browsers support JavaScript modules, loading them individually would result
in aslow startup due to many network requests.

To solve this, tools like `esbuild` bundle these files together into larger JavaScript files.
However, making one huge bundle is also bad because it takes a long time to download, parse,
and execute. On mobile devices this can be a huge hinderance to the user experience.

A smarter approach is splitting the bundle using **dynamic imports**, which allows the browser
to load parts of the app only when needed.

 A **clean, one-way dependency graph** leads to better performance by keeping bundles smaller
 and more efficient. That is why it is important to understand how to structure our code efficiently.

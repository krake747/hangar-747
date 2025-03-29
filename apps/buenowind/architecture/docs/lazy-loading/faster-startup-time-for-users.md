# Faster Startup Time for Users

With lazy loading, we only load the parts of the app that are necessary at the start, instead of
loading everything all at once. This allows the app to load faster, as we're downloading less code initially.

For instance, if the app has many features, we might only load the main part of the app and the first
feature the user needs. The more features we have, the less code needs to be downloaded up front.

Even though modern internet connections and caching can help, lazy loading still offers an advantage.
After the code is downloaded, the app needs to "read" and execute it, which can take time.
Lazy loading keeps the initial code smaller, allowing the app to run faster, even on slower devices.

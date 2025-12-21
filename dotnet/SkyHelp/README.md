# SkyHelp

Starting with .NET 10 we can have support for file-based apps.

Key benefits:

- Reduced boilerplate for simple applications
- Self-contained source files
- Native AOT publishing enabled by default

## Run the application

```bash
dotnet run main.cs
```

## Publishing Options

**Native AOT (default)**

- Pros: extremely fast startup, low memory, single native binary
- Cons: requires `clang`/`gcc` in build image, bigger build, harder to debug
- Best for production microservices where startup matters

**Normal IL-based publish (`/p:PublishAot=false`)**

- Pros: simpler Docker builds, no extra compiler needed, easier development and CI
- Cons: slower startup (~200–400ms), larger memory footprint, requires .NET runtime
- Good for development, testing, and CI/CD pipelines

Example:

```bash
dotnet publish main.cs -c Release -o /app /p:UseAppHost=false /p:PublishAot=false
```

## Additional Resources

[File based apps](https://learn.microsoft.com/en-us/dotnet/core/sdk/file-based-apps)
[Tutorial to build](https://learn.microsoft.com/en-us/dotnet/csharp/fundamentals/tutorials/file-based-programs)

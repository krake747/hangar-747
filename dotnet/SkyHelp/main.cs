#:sdk Microsoft.NET.Sdk.Web

var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

app.MapGet("/", () => Results.Ok(new { message = "Hello from SkyHelp!" }));

app.Run();

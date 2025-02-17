var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

app.MapGet("/", () => "SkyOps! How may we help you?");

app.Run();

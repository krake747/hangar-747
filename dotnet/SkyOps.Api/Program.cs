var builder = WebApplication.CreateBuilder(args);

builder.Services.AddHttpClient();

var app = builder.Build();

app.MapGet(
    "/",
    static () => Results.Ok(new { Message = "SkyOps! How may we help you?" })
);

app.Run();

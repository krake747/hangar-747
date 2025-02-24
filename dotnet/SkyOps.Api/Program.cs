using System.Text.Json;
using Microsoft.AspNetCore.Mvc;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddHttpClient();

var app = builder.Build();

app.MapGet(
    "/",
    static () => Results.Ok(new { Message = "SkyOps! How may we help you?" })
);

app.MapGet(
    "/skyconvert",
    static async (
        [FromServices] ILogger<Program> logger, [FromServices] IHttpClientFactory factory,
        [FromServices] IConfiguration config
    ) =>
    {
        logger.LogInformation("Processing SkyOps to SkyConvert request");
        using var client = factory.CreateClient();
        var skyConvertUrl = config.GetValue<string>("Networking:SkyConvert") ?? "";
        client.BaseAddress = new Uri(skyConvertUrl);
        var response = await client.GetFromJsonAsync<JsonElement>("");
        return Results.Ok(new { Message = response.GetProperty("message").GetString() });
    }
);

app.Run();

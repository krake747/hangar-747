using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.FileProviders;
using SkyConvert.Api.Common;
using SkyConvert.Api.Networking;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddSingleton(static sp =>
{
    var config = sp.GetRequiredService<IConfiguration>();
    return new VolumeManager(config, Directory.GetCurrentDirectory());
});


builder.Services.AddHttpClient<IGotenbergClient, GotenbergClient>("gotenberg", static (sp, client) =>
{
    var connectionString = sp.GetRequiredService<IConfiguration>().GetConnectionString("Gotenberg") ?? "";
    client.BaseAddress = new Uri(connectionString);
});

var app = builder.Build();

app.UseStaticFiles(new StaticFileOptions
{
    FileProvider = new PhysicalFileProvider(app.Services.GetRequiredService<VolumeManager>().Thumbnails.FullName),
    RequestPath = "/thumbnails"
});

app.MapGet(
    "/",
    () => Results.Ok(new { Message = "Hello from SkyConvert! Would you like to convert a document?" })
);

app.MapGet(
    "/health/gotenberg",
    static async ([FromServices] IHttpClientFactory factory, CancellationToken token = default) =>
    {
        using var client = factory.CreateClient("gotenberg");
        return Results.Ok(await client.GetAsync("health", token));
    }
);

app.MapPost(
    "thumbnails",
    static async (
        [FromServices] ILogger<Program> logger, [FromServices] IGotenbergClient gotenberg,
        [FromServices] VolumeManager directoryManager, HttpRequest request, CancellationToken token = default
    ) =>
    {
        logger.LogInformation("Processing generate thumbnail request");

        if (request.HasFormContentType is false)
        {
            return Results.BadRequest("No files were uploaded.");
        }

        if (request.Form.Files.Count is 0)
        {
            return Results.BadRequest("No files were uploaded.");
        }

        var thumbnail = await gotenberg.GenerateThumbnails(request, directoryManager.Thumbnails, token);

        string[] thumbnails = [thumbnail];

        return Results.Ok(new { Thumbnails = thumbnails });
    }
);

app.MapGet(
    "thumbnails",
    static ([FromServices] VolumeManager directoryManager) =>
    {
        string[] thumbnails =
        [
            ..Directory
                .EnumerateFiles(directoryManager.Thumbnails.FullName)
                .Select(f => $"/thumbnails/{Path.GetFileName(f)}")
        ];

        return Results.Ok(new { Thumbnails = thumbnails });
    }
);

app.Run();

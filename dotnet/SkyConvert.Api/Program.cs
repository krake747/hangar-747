using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.FileProviders;
using SkyConvert.Api.Common;
using SkyConvert.Api.Networking;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddHttpClient<IGotenbergClient, GotenbergClient>("gotenberg", static (sp, client) =>
{
    var connectionString = sp.GetRequiredService<IConfiguration>().GetConnectionString("Gotenberg") ?? "";
    client.BaseAddress = new Uri(connectionString);
});

var thumbnailsDirectory = Directory.GetCurrentDirectory()
    .Pipe(root => Path.Combine(root, builder.Configuration.GetValue<string>("Directory:Thumbnails") ?? "thumbnails"))
    .Pipe(Directory.CreateDirectory);

var app = builder.Build();

app.UseStaticFiles(new StaticFileOptions
{
    FileProvider = new PhysicalFileProvider(thumbnailsDirectory.FullName),
    RequestPath = "/thumbnails"
});

app.MapGet(
    "/",
    () => Results.Ok(new { Message = "Hello from SkyConvert! Would you like to convert a document?" })
);

app.MapGet(
    "/health/gotenberg",
    async ([FromServices] IHttpClientFactory factory, CancellationToken token = default) =>
    {
        using var client = factory.CreateClient("gotenberg");
        return Results.Ok(await client.GetAsync("health", token));
    }
);

app.MapPost(
    "thumbnails",
    static async (
        [FromServices] ILogger<Program> logger, [FromServices] GotenbergClient gotenberg,
        HttpRequest request, CancellationToken token = default
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

        var t = await gotenberg.GenerateThumbnails(request, token);

        string[] thumbnails = [t];
        return Results.Ok(new { Thumbnails = thumbnails });
    }
);

app.Run();

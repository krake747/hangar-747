var builder = WebApplication.CreateBuilder(args);

builder.Services.AddHttpClient();

var app = builder.Build();

app.MapGet(
    "/",
    static () => Results.Ok(
        new
        {
            Message = "SkyOps! How may we help you?",
            RedirectToSkyHelp = "/redirect/skyhelp"
        }
    )
);

app.MapGet(
    "/redirect/skyhelp",
    async static (HttpClient httpClient, IConfiguration config) =>
    {
        var skyHelpUrl = config["Networking:SkyHelp"] ?? "http://skyhelp:8080/";
        var response = await httpClient.GetAsync(skyHelpUrl);
        var content = await response.Content.ReadAsStringAsync();
        return Results.Content(content, "application/json");
    }
);

app.Run();

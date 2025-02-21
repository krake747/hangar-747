var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello from SkyConvert! Would you like to convert a document?");

app.Run();

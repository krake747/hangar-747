using System.Net.Http.Headers;
using ImageMagick;

namespace SkyConvert.Api.Networking;

internal interface IGotenbergClient
{
    Task<string> GenerateThumbnails(HttpRequest request, CancellationToken token = default);
}

internal sealed class GotenbergClient(HttpClient httpClient) : IGotenbergClient
{
    public async Task<string> GenerateThumbnails(HttpRequest request, CancellationToken token = default)
    {
        var formFile = request.Form.Files[0];

        if (Path.GetExtension(formFile.FileName) is ".jpg" or ".jpeg" or ".png" or ".gif")
        {
            return await ConvertImageFileAsync(formFile, token);
        }

        using var multipartFormDataContent = new MultipartFormDataContent();
        using var memoryStream = new MemoryStream();
        await formFile.CopyToAsync(memoryStream, token);
        memoryStream.Seek(0, SeekOrigin.Begin);

        var fileContent = new StreamContent(memoryStream);
        fileContent.Headers.ContentType = MediaTypeHeaderValue.Parse("application/octet-stream");
        multipartFormDataContent.Add(new StringContent("1-1"), "nativePageRanges");
        multipartFormDataContent.Add(fileContent, "files", formFile.FileName);

        var response = await httpClient.PostAsync("/forms/libreoffice/convert", multipartFormDataContent, token);
        if (response.IsSuccessStatusCode is false)
        {
            throw new Exception("Failed to convert the file to PDF.");
        }

        var pdfFile = Path.ChangeExtension(formFile.FileName, ".pdf");
        await using var fileStream = new FileStream(pdfFile, FileMode.Create);
        await response.Content.CopyToAsync(fileStream, token);

        return pdfFile;
    }

    private static async Task<string> ConvertImageFileAsync(IFormFile formFile, CancellationToken token = default)
    {
        using var memoryStream = new MemoryStream();
        await formFile.CopyToAsync(memoryStream, token);
        memoryStream.Seek(0, SeekOrigin.Begin);

        var magickImage = new MagickImage(memoryStream)
        {
            Format = MagickFormat.WebP
        };

        var thumbnailFile = Path.ChangeExtension(formFile.FileName, ".webp");

        await magickImage.WriteAsync(thumbnailFile, token);

        return thumbnailFile;
    }
}

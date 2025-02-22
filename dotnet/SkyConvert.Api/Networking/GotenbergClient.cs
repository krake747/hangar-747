using System.Net.Http.Headers;
using ImageMagick;

namespace SkyConvert.Api.Networking;

internal interface IGotenbergClient
{
    Task<string> GenerateThumbnails(HttpRequest request, DirectoryInfo thumbnailsDirectory,
        CancellationToken token = default);
}

internal sealed class GotenbergClient(HttpClient httpClient) : IGotenbergClient
{
    public async Task<string> GenerateThumbnails(HttpRequest request, DirectoryInfo thumbnailsDirectory,
        CancellationToken token = default)
    {
        var formFile = request.Form.Files[0];

        if (Path.GetExtension(formFile.FileName) is ".jpg" or ".jpeg" or ".png")
        {
            using var memoryStream = new MemoryStream();
            await formFile.CopyToAsync(memoryStream, token);
            memoryStream.Seek(0, SeekOrigin.Begin);

            var webImage = await ConvertToWebpAsync(memoryStream, formFile.FileName, token);
            return webImage;
        }
        else
        {
            using var multipartFormDataContent = new MultipartFormDataContent();
            using var memoryStream = new MemoryStream();
            await formFile.CopyToAsync(memoryStream, token);
            memoryStream.Seek(0, SeekOrigin.Begin);

            using var fileContent = new StreamContent(memoryStream);
            fileContent.Headers.ContentType = MediaTypeHeaderValue.Parse("application/octet-stream");

            multipartFormDataContent.Add(new StringContent("1-1"), "nativePageRanges");
            multipartFormDataContent.Add(fileContent, "files", formFile.FileName);

            var response = await httpClient.PostAsync("/forms/libreoffice/convert", multipartFormDataContent, token);
            if (response.IsSuccessStatusCode is false)
            {
                throw new Exception("Failed to convert the file to PDF.");
            }

            var pdf = Path.Combine(thumbnailsDirectory.FullName, Path.ChangeExtension(formFile.FileName, ".pdf"));
            // await using var fileStream = new FileStream(pdf, FileMode.Create);
            await using var pdfStream = await response.Content.ReadAsStreamAsync(token);
            // await response.Content.CopyToAsync(fileStream, token);

            var webpImage = await ConvertToWebpAsync(pdfStream, pdf, token);
            return webpImage;
        }
    }

    private static async Task<string> ConvertToWebpAsync(Stream pdfStream, string fileName,
        CancellationToken token = default)
    {
        var magickImage = new MagickImage(pdfStream)
        {
            Format = MagickFormat.WebP
        };

        var webpImage = Path.ChangeExtension(fileName, ".webp");

        await magickImage.WriteAsync(webpImage, token);

        return webpImage;
    }
}

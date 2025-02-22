namespace SkyConvert.Api.Common;

internal sealed class VolumeManager(IConfiguration config, string rootPath)
{
    public DirectoryInfo Thumbnails { get; } =
        Directory.CreateDirectory(Path.Combine(rootPath, config["Volumes:Thumbnails"]!));
}

FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
WORKDIR /src

# Install Native AOT prerequisites
RUN apt-get update && apt-get install -y clang gcc libc6-dev

COPY SkyHelp/main.cs ./main.cs

RUN dotnet publish main.cs \
    -c Release \
    -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:10.0 AS base
WORKDIR /app
COPY --from=build /app/publish .
EXPOSE 8080
ENTRYPOINT ["dotnet", "main.dll"]

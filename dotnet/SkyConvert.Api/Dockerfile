FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["SkyConvert.Api/SkyConvert.Api.csproj", "SkyConvert.Api/"]
RUN dotnet restore "SkyConvert.Api/SkyConvert.Api.csproj"
COPY . .
WORKDIR "/src/SkyConvert.Api"
RUN dotnet publish "SkyConvert.Api.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base
WORKDIR /app
RUN apt-get -y update && apt-get -y install ghostscript
EXPOSE 8080
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "SkyConvert.Api.dll"]

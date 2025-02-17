FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["SkyOps.Api/SkyOps.Api.csproj", "SkyOps.Api/"]
RUN dotnet restore "SkyOps.Api/SkyOps.Api.csproj"
COPY . .
WORKDIR "/src/SkyOps.Api"
RUN dotnet publish "SkyOps.Api.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base
WORKDIR /app
EXPOSE 8080
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "SkyOps.Api.dll"]
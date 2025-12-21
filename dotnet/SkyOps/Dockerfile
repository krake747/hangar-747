FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["SkyOps/SkyOps.csproj", "SkyOps/"]
RUN dotnet restore "SkyOps/SkyOps.csproj"
COPY . .
WORKDIR "/src/SkyOps"
RUN dotnet publish "SkyOps.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM mcr.microsoft.com/dotnet/aspnet:10.0 AS base
WORKDIR /app
EXPOSE 8080
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "SkyOps.dll"]

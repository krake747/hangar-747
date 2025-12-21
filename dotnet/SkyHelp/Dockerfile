FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
WORKDIR /src

# Native AOT prerequisites
# RUN apt-get update && apt-get install -y clang gcc libc6-dev

COPY SkyHelp/main.cs ./main.cs

RUN dotnet publish main.cs \
    -c Release \
    -r linux-x64 \
    -o /app/publish \
    -p:PublishAot=false

FROM mcr.microsoft.com/dotnet/aspnet:10.0 AS base
WORKDIR /app
COPY --from=build /app/publish .
EXPOSE 8080
ENTRYPOINT ["./main"]

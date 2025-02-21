# SkyConvert.Api

Sky Operations to coordinate flights and schedules.

## Setup

```shell
dotnet new web -n SkyConvert.Api
dotnet sln Hangar747.sln add SkyConvert.Api/SkyConvert.Api.csproj
```

## Local Dev

Execute from root directory the following command to run the api:

```shell
dotnet run --project SkyConvert.Api
```

Test the API:

```shell
curl http://localhost:5034
```

## Docker Image

Execute from root directory the following command to build Docker image:

```docker
docker build -t skyconvert-api:latest -f SkyConvert.Api/Dockerfile .
```

We can start a new docker container via

```docker
docker run -d -p 7000:8080 skyconvert-api:latest
```

Here, we check if we can communicate with the Api running inside of the Docker container

```shell
curl http://localhost:7000
```

## Minikube

If we want to load the local Docker image into Minikube we run

```shell
minikube image load skyconvert-api:latest
```

# SkyOps.Api

Sky Operations to coordinate flights and schedules.

## Setup

```shell
dotnet new web -n SkyOps.Api
dotnet sln Hangar747.sln add SkyOps.Api/SkyOps.Api.csproj
```

## Local Dev

Execute from root directory the following command to run the api:

```shell
dotnet run --project SkyOps.Api
```

Test the API:

```shell
curl http://localhost:5134
```

## Docker Image

Execute from root directory the following command to build Docker image:

```docker
docker build -t skyops-api:latest -f SkyOps.Api/Dockerfile .
```

We can start a new docker container via

```docker
docker run -d -p 7000:8080 skyops-api:latest
```

Here, we check if we can communicate with the Api running inside of the Docker container

```shell
curl http://localhost:7000
```

## Minikube

If we want to load the local Docker image into Minikube we run

```shell
minikube image load skyops-api:latest
```

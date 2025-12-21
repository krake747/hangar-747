# SkyOps

Sky Operations to coordinate flights and schedules.

## Setup

```bash
dotnet new web -n SkyOps
dotnet sln Hangar747.slnx add SkyOps/SkyOps.csproj
```

## Local Dev

Execute from the root directory the following command to run the api:

```bash
dotnet run --project SkyOps
```

Test the API:

```bash
curl http://localhost:5134
```

## Docker Image

Execute from the root directory the following command to build a Docker image:

```docker
docker build -t skyops:latest -f SkyOps/Dockerfile .
```

We can start a new docker container via

```docker
docker run -d -p 7000:8080 skyops:latest
```

Here, we check if we can communicate with the Api running inside the Docker container

```bash
curl http://localhost:7000
```

## Minikube

If we want to load the local Docker image into Minikube, we run

```bash
minikube image load skyops:latest
```

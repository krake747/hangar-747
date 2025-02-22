# SkyConvert.Api

Sky Conversions to convert documents.

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

### User Secrets

We use .NET's User Secrets to manage the `Gotenberg` connection string for local development.

```json
{
    "ConnectionStrings": {
        "Gotenberg": "http://localhost:3000/"
    }
}
```

To check if `Gotenberg` is up and running, we can check via

```shell
curl http://localhost:5034/health/gotenberg
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

Once we apply the `k8s/skyconvert.yaml` deployment the pod should start up.

Let's say we want to inspect the `thumbnails-volume`.

```shell
k get pods -l app=skyconvert -n=hangar-747
```

We can access and inspect the pod's mounted volume with `Bash`.

```shell
k exec -it skyconvert-[...] -- /bin/bash
```

We inspect it with

```shell
ls -lah /thumbnails
```

If we want to monitor changes in real time:

```shell
watch ls -lah /thumbnails
```

**Note:** If we get the `watch` command not found then we can try and install it with

```shell
apt update && apt install -y procps
```

Then we try to re-run the watch command above and create a test file.

```shell
echo "hello test 2" > /thumbnails/test2.txt
```

## Gotenberg

It can be that we need `Ghostscript` for converting Pdfs to images.

```shell
sudo apt update && sudo apt install ghostscript
```

Testing `Gotenberg` locally.

```shell
curl -X POST -F "files=@.input/x.txt" -F "nativePageRanges=1-1" -o ".output/x.pdf" http://localhost:3000/forms/libreoffice/convert
```

Testing `SkyConvert` locally from solution root

```shell
curl -X POST -F "files=@.input/x.txt" http://localhost:5034/thumbnails
```

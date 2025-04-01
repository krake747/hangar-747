# Docs

To generate the documentation website we just need to run from the root folder which contains the `compose.yml`
and `compose.override.yaml`.

```docker
docker compose watch buenowind.docs
```

Using `docker` we can avoid installing Python 🐍 on our local machine.

For more info on customization visit [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/)

## Build

```shell
cd architecture
docker build -f Dockerfile -t my-mkdocs .
docker run -p 8000:8000 my-mkdocs
```

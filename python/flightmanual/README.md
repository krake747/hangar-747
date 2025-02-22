# Flight Manual

The flight manual is static website to host my documentation of Hanager 747. I am using [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/) because it is easy to setup and use. 

## Quick start

First we need to install `Python`:

```shell
sudo apt install python3
sudo apt install python3.12-venv
```

We then create a virtual environment and activate it from the project root directory

```shell
python3 -m venv .venv
. .venv/bin/activate
```

**Note**: The terminal can show when the virtual environment is active.

We create a `requirements.txt` file which contains `mkdocs-material`

```shell
echo "mkdocs-material" > requirements.txt
```

Now that we have a `requirements.txt` we can setup `Material for MkDocs`.

```shell
pip install -r requirements.txt
```

To upgrade `mkdocs-material` we can run 

```shell
pip install --upgrade --force-reinstall mkdocs-material
pip show mkdocs-material
```

## Useful Commands

- `mkdocs new [dir-name]` - Creates a new project.
- `mkdocs serve` - Starts the live-reloading docs server.
- `mkdocs build` - Builds the documentation site.
- `mkdocs -h` - Prints help message and exit.

[![Built with Material for MkDocs](https://img.shields.io/badge/Material_for_MkDocs-526CFE?style=for-the-badge&logo=MaterialForMkDocs&logoColor=white)](https://squidfunk.github.io/mkdocs-material/)

and [MkDocs](https://www.mkdocs.org).


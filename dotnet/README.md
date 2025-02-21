# .NET Playground

## Install .NET on Linux

On WSL I am using the Ubuntu 24.04 distro and to install .NET I followed the steps highlighted in this [page](https://learn.microsoft.com/en-us/dotnet/core/install/linux-ubuntu-install?pivots=os-linux-ubuntu-2404&tabs=dotnet9).

```shell
sudo add-apt-repository ppa:dotnet/backports
```

```shell
sudo apt-get update && sudo apt-get install -y dotnet-sdk-9.0
```

## Setup the Solution File

```shell
dotnet new sln -n Hangar747
```

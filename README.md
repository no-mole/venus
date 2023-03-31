# <img src="static/icon.svg" width="33px" height="33px">   VENUS

**Note**: The `main` branch may be in an *unstable or even broken state* during development. For stable versions, see [releases][github-release].

Venus is a distributed, highly available configuration management & service registration & service Discovery distributed infrastructure.

Venus provides several key features:
* **Service Registration/Discovery** - Venus makes it simple for services to register,
clients in various languages can quickly register service portals,
lease mechanism can invalidate the portal after the service is offline.
* **Key/Value Storage** - A flexible key/value store enables storing
  dynamic configuration. The simple HTTP API/GRPC API/CommandLine Tools makes it easy to use anywhere.
* **Highly Available** - Venus uses raft protocol to ensure data synchronization to enough nodes.
* **Distributed** - Venus uses GRPC and HTTP/HTTPS to provide service portals, uses GRPC to process data communication of each node, 
and uses **boltdb** to store raft log and persistent data.
* **UI** - Venus comes with a panel and support password login or oidc login，it is very convenient to manage configuration items and services.

Venus using go development, cross-platform deployment is possible.

Quick Start
---
Install binary
```shell
go install github.com/no-mole/venus/agent@latest
```
View Help
```shell
agent -h
```
Start single node
```shell
agent --boot
```

Install command line tools
```shell
go install github.com/no-mole/venus/vtl@latest
```
View Help
```shell
vtl -h
```

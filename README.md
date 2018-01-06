[![Datalayer](http://datalayer.io/img/logo-datalayer-horizontal.png)](http://datalayer.io)

This folder contains the `Golang` REST and WebSocket server for `Kuber`.

`git clone` this repository in your `go` path.

## Hack Source Code

Start the `K8S Proxy`.

```
kubectl proxy
```

Start the `Kuber Server`.

```console
cd $GOPATH/src/github.com/datalayer/kuber
go run main.go server --apiserver-host=http://localhost:8001
```

We ship a snapshot of the user interface (use the [kuber-plane](https://github.com/datalayer/kuber-plane) repository for the latest version).

You can now browse [http://localhost:9091](http://localhost:9091) or make REST call to the API.

## Manage Dependencies

This repository ships the `vendor` dependencies to ensure comptability.

If you want to get your own dependencies, use the `dep` tool.

```console
dep init
dep ensure
```

## Build Binary

```console
go build
```

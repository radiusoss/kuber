[![Datalayer](http://datalayer.io/img/logo-datalayer-horizontal.png)](http://datalayer.io)

This folder contains the `Golang` REST and WebSocket server for `Kuber`.

`git clone` this repository in your `go` path.

## Hack the Code

Start the `K8S Proxy`.

```
kubectl proxy
```

Start the `Kuber Server`.

```
cd $GOPATH/src/github.com/datalayer/kuber
go run main.go server --apiserver-host=http://localhost:8001
```

## Manage the Dependencies

This repository ships the `vendor` dependencies to ensure comptability.

If you want to get your own dependencies, use the `dep` tool.

```shell
dep init
dep ensure
```

## Build a Binary

```shell
go build
```

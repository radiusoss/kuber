[![Datalayer](http://datalayer.io/img/logo-datalayer-horizontal.png)](http://datalayer.io)

This folder contains the `Golang` source code for `Kuber`, a tool to create and operate Kubernetes clusters.

It ships a REST and WebSocket server and a CLI (Command Line Interface).

## Usage

`git clone` this repository in your `go` path.

Start the `K8S Proxy`.

```console
kubectl proxy
```

Start the `Kuber Server`.

```console
cd $GOPATH/src/github.com/datalayer/kuber
go run main.go server --apiserver-host=http://localhost:8001
```

We ship a (probably outdated) snapshot of the user interface (use the [Kuber Plane](https://github.com/datalayer/kuber-plane) repository for the latest version).

You can now browse [http://localhost:9091](http://localhost:9091) or make REST call to the API.

## Build Binary

```console
cd $GOPATH/src/github.com/datalayer/kuber
go build
```

## Dependencies

This repository ships the `vendor` dependencies to ensure comptability.

If you want to get your own dependencies, use the `dep` tool.

```console
dep init
dep ensure
```

## K8S Cluster

From your Linux laptop with [Helm](https://github.com/kubernetes/helm/releases) available, run the following.

```shell
export AWS_ACCESS_KEY_ID=<your-aws-key-id>
export AWS_SECRET_ACCESS_KEY=<your-aws-key-secret>
kuber create kuber -p aws
kuber apply kuber -v 4
```

Check the cluster is running.

```console
watch kubectl get nodes; watch kubectl get pods --all-namespaces; kubectl proxy
```

Delete the cluster.

```console
kuber delete kuber -v 4 --purge
```

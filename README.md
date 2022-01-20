# [containerd](https://github.com/containerd/containerd) darwin snapshotter plugin

[![CI](https://github.com/ukontainer/darwin-snapshotter/actions/workflows/ci.yml/badge.svg)](https://github.com/ukontainer/darwin-snapshotter/actions/workflows/ci.yml)

Darwin snapshotter is a snapshotter implementation which works on darwin platform.  It is backed with sparsebundle images (.sparsebundle) which is created from each layer.  Due to lack of public interface (or API) to create sparsebundle images on macOS, each image creation and mount operation invoke hdiutil command, which is unfortunatelly costly (in terms of both storage space required and time completing each layer operation).

But with this snapshotter, the mount package of containerd doesn't depend on bind mount features, which is transparent from other mount implementations.

## Setup

```console
$ go build ./cmd/containerd-darwin-snapshotter-grpc
$ go build ./cmd/mount_containerd_darwin
```


## Usage

```console
$ ./containerd-darwin-snapshotter-grpc /var/run/containerd-darwin-snapshotter.sock /var/lib/containerd/darwin
```

containerd has to be also configured to utilize this snapshotter.  Below is an example configuration to add the snapshotter.

```
[proxy_plugins]
  [proxy_plugins.darwin]
    type = "snapshot"
    address = "/tmp/ctrd/run/containerd-darwin-snapshotter.sock"
```

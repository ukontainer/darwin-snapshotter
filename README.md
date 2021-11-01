# [containerd](https://github.com/containerd/containerd) darwin snapshotter plugin

[![Build Status]()

Darwin snapshotter plugin for containerd.

## Setup

```console
$ go build ./cmd/darwin-snapshotter-grpc
```

Note that the daemon binary needs to be exactly the version used for building the shared library.

Please refer to [`Makefile`](./Makefile) for the version known to work with.

## Usage

```console
$ ./darwin-snapshotter-grpc /tmp/ctrd/run/containerd-darwin-snapshotter.sock /tmp/ctrd/var/lib/containerd/darwin
```

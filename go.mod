module github.com/ukontainer/darwin-snapshotter

go 1.17

replace github.com/containerd/containerd => github.com/ukontainer/containerd v1.5.1-0.20211025224442-b6810835c126

require (
	github.com/containerd/containerd v1.5.7
	github.com/containerd/containerd/api v1.6.0-beta.1
	github.com/containerd/continuity v0.2.1
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.41.0
)

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/Microsoft/hcsshim v0.9.0 // indirect
	github.com/containerd/cgroups v1.0.2 // indirect
	github.com/containerd/ttrpc v1.0.2 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/moby/sys/mountinfo v0.4.2-0.20211022201527-95edfa939201 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2-0.20210819154149-5ad6f50d6283 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210915083310-ed5796bab164 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gotest.tools/v3 v3.0.3 // indirect
)
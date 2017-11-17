# dockerd
Docker daemon api


## Configure in your docker daemon
on Linux platform, you can update docker daemon as below:
```sh
cd /etc/systemd/system/docker.service.d/
touch tcp.conf
echo
"[Service]
ExecStart=
ExecStart=/usr/bin/docker daemon -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock"
>> tcp.conf
systemctl daemon-reload
systemctl restart docker

```

## Clone source
```sh
git clone https://github.com/nightlegend/dockerd.git

```

## Get package and build

```sh
go get
go build

```

## Start

```sh
go run main.go

```

## Main function

```go
package main

import (
	"github.com/docker/docker/api/types/container"
	new_client "github.com/nightlegend/dockerd/client"
	"github.com/nightlegend/dockerd/containers"
	"github.com/nightlegend/dockerd/images"
)

func main() {
  // New a docker client object.
	cli, ctx := new_client.NewClinet("http://localhost:2375")
  // pull a docker images, if registry is private,you can use full name(eg: xxx.xxx.com/os/centos:latest)
	images.PullImage(ctx, cli, "busybox")
  // create a container and run it.
	containers.RunContainer(ctx, cli, &container.Config{
		Image: "busybox",
		Cmd:   []string{"echo", "hello world"},
	})
	containers.ShowContainerLogs(ctx, cli, "b64726d90208")
	images.ListAllImage(ctx, cli)
	containers.RestartContainer(ctx, cli, "b64726d90208", nil)
	containers.ContainerStats(ctx, cli, "b64726d90208")
	containers.InspectContainer(ctx, cli, "b64726d90208")
	containers.StopContainer(ctx, cli, "b64726d90208")
	containers.StartContainer(ctx, cli, "b64726d90208")
}

```

package main

import (
	"github.com/docker/docker/api/types/container"
	new_client "github.com/nightlegend/dockerd/client"
	"github.com/nightlegend/dockerd/containers"
	"github.com/nightlegend/dockerd/images"
)

func main() {
	cli, ctx := new_client.NewClinet("http://localhost:2375")
	images.PullImage(ctx, cli, "busybox")
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

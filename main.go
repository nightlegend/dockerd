package main

import (
	new_client "github.com/nightlegend/dockerd/client"
	"github.com/nightlegend/dockerd/containers"
	"github.com/nightlegend/dockerd/images"
)

func main() {
	cli, ctx := new_client.NewClinet("http://localhost:2375")
	// images.ImagePull(ctx, cli, "intdocker.cargosmart.com/os/busybox:latest")
	// containers.ContainerRun(ctx, cli, &container.Config{
	// Image: "busybox",
	// Cmd:   []string{"echo", "hello world"},
	// })
	// containers.ContainerShowLogs(ctx, cli, "66cb79d0abe4")
	// images.ImageListAll(ctx, cli)
	containers.ContainerRestart(ctx, cli, "66cb79d0abe4", nil)
	// containers.ContainerStats(ctx, cli, "66cb79d0abe4")
	// containers.ContainerInspect(ctx, cli, "66cb79d0abe4")
	// containers.ContainerStop(ctx, cli, "66cb79d0abe4")
	// containers.ContainerStart(ctx, cli, "66cb79d0abe4")
	// images.ImageTag(ctx, cli, "6ad733544a63", "intdocker.cargosmart.com/os/busybox:dev1.0")
	images.ImageInfo(ctx, cli, "6ad733544a63")
	images.ImageBuild(ctx, cli)
	images.ImageClean(ctx, cli, "ce5e58f82f64")
}

package containers

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

// RunContainer and start a container
func RunContainer(ctx context.Context, cli *client.Client, cfg *container.Config) {
	resp, err := cli.ContainerCreate(ctx, cfg, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	log.Println(resp.ID)
}

// StopContainer is stop a running container.
func StopContainer(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		panic(err)
	}
}

// StartContainer is start a stop container.
func StartContainer(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}

// ShowContainerLogs is show a container logs.
func ShowContainerLogs(ctx context.Context, cli *client.Client, containerID string) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

// RestartContainer is restart a running container.
func RestartContainer(ctx context.Context, cli *client.Client, containerID string, timeout *time.Duration) {
	err := cli.ContainerRestart(ctx, containerID, timeout)
	if err != nil {
		panic(err)
	}
}

// ContainerStats is get a container stats.
func ContainerStats(ctx context.Context, cli *client.Client, containerID string) {
	stats, err := cli.ContainerStats(ctx, containerID, true)
	if err != nil {
		panic(err)
	}
	log.Println(stats.Body)
	log.Printf("result is :%s", stats.Body)
}

// InspectContainer is inspect a container detail information.
func InspectContainer(ctx context.Context, cli *client.Client, containerID string) {
	result, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	// var m types.ContainerJSON
	log.Println(result.Config)
}

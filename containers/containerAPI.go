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

// ContainerRun and start a container
func ContainerRun(ctx context.Context, cli *client.Client, cfg *container.Config) {
	resp, err := cli.ContainerCreate(ctx, cfg, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	log.Println(resp.ID)
}

// ContainerStop is stop a running container.
func ContainerStop(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		panic(err)
	}
}

// ContainerStart is start a stop container.
func ContainerStart(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}

// ContainerShowLogs is show a container logs.
func ContainerShowLogs(ctx context.Context, cli *client.Client, containerID string) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

// ContainerRestart is restart a running container.
func ContainerRestart(ctx context.Context, cli *client.Client, containerID string, timeout *time.Duration) {
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

// ContainerInspect is inspect a container detail information.
func ContainerInspect(ctx context.Context, cli *client.Client, containerID string) {
	result, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	// var m types.ContainerJSON
	log.Println(result.Config)
}

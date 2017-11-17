package images

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

// PullImage is clinet pull a docker images.
func PullImage(ctx context.Context, cli *client.Client, name string) {
	out, err := cli.ImagePull(ctx, name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

// ListAllImage is list all images.
func ListAllImage(ctx context.Context, cli *client.Client) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}

package client

import (

	// "dockerd/images"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

//NewClinet create a new client object.
func NewClinet(url string) (*client.Client, context.Context) {
	ctx := context.Background()
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient(url, "v1.23", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	return cli, ctx
}

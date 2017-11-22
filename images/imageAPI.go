package images

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"archive/tar"
	"bytes"

	"io/ioutil"
)

// ImagePull is clinet pull a docker images.
func ImagePull(ctx context.Context, cli *client.Client, name string) {
	out, err := cli.ImagePull(ctx, name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

// ImageListAll is list all images.
func ImageListAll(ctx context.Context, cli *client.Client) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}

// ImageTag is give a new tag to your image.
func ImageTag(ctx context.Context, cli *client.Client, imageID string, ref string) {
	err := cli.ImageTag(ctx, imageID, ref)
	if err != nil {
		log.Println(err)
	}
}

// ImageInfo is inspect a image detail information.
func ImageInfo(ctx context.Context, cli *client.Client, imageID string) {
	imageInspectInfo, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		log.Println(err)
	}
	log.Println(imageInspectInfo)
}

// ImageBuild is build a image by yourself.
func ImageBuild(ctx context.Context, cli *client.Client) {
	dockerBuildContext, err := os.Open("D:/DT/cs_acz_si.tar.gz")
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := "myDockerfile"
	// the Dockerfile path is your local location.
	dockerFileReader, err := os.Open("D:/DT/Dockerfile")
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Tags:       []string{"test_images:v1"},
			Context:    dockerBuildContext,
			Dockerfile: dockerFile,
			Remove:     true})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}

// ImageClean is delete a unused image.
func ImageClean(ctx context.Context, cli *client.Client, imageID string) {
	cleanResult, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force: true,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(cleanResult)
}

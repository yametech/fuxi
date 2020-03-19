package client

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
)

func TestDockerCli_PullImage(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	err = cli.PullImage(context.Background(), "nginx")
	if err != nil {
		t.Error(err)
	}
}

func TestDockerCli_RemoveImage(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	err = cli.RemoveImage(context.Background(), "nginx")
	if err != nil {
		t.Error(err)
	}
}

func TestDockerCli_BuildImage(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	err = cli.BuildImage(context.Background(), types.ImageBuildOptions{})
	if err != nil {
		t.Error(err)
	}
}

func TestDockerCli_ListImages(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	images, err := cli.ListImages(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, value := range images {
		fmt.Println(value.ID)
		fmt.Println(value.Containers)
		fmt.Println(value.ParentID)
		fmt.Println(value.Created)
		fmt.Println(value.Labels)
		fmt.Println(value.RepoTags)
	}
}

func TestDockerCli_InspectImage(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.InspectImage(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_PushImage(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	auth := types.AuthConfig{
		Username:      "ym",
		Password:      "123123",
		Auth:          "",
		Email:         "",
		ServerAddress: "",
		IdentityToken: "",
		RegistryToken: "",
	}
	if err = cli.RegistryLogin(context.Background(), auth); err != nil {
		t.Error(err)
	}
	if err = cli.PushImage(context.Background(), "test/nginx"); err != nil {
		t.Error(err)
	}

}

func TestDockerCli_CreateContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	err = cli.CreateContainer(context.Background(), "nginx", nil, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerCli_ContainerClean(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.ContainerClean(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}

}

func TestDockerCli_StartContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.StartContainer(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_StopContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.StopContainer(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_RestartContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.RestartContainer(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_KillContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.KillContainer(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_PauseContainer(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	if err = cli.PauseContainer(context.Background(), "nginx"); err != nil {
		t.Error(err)
	}
}

func TestDockerCli_GetContainerLogs(t *testing.T) {
	cli, err := NewDefaultDockerClient()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()
	res, err := cli.GetContainerLogs(context.Background(), "nginx", types.ContainerLogsOptions{})
	if err != nil {
		t.Error(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res)
	if err != nil {
		t.Error(err)
	}
	s := buf.String()
	t.Log(s)
}

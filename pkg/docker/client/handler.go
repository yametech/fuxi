package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/docker/go-connections/tlsconfig"

	containertypes "github.com/docker/docker/api/types/container"
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"

	"github.com/yametech/fuxi/pkg/docker/inspect"

	dockerclient "github.com/docker/docker/client"
	"github.com/docker/docker/opts"

	"github.com/yametech/fuxi/pkg/docker/stream"

	"github.com/docker/docker/pkg/term"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

//see doc  in  https://docs.docker.com/engine/api/v1.40/
//ImageHandler  image handler
type ImageHandler interface {
	PullImage(context context.Context, name string) error
	RemoveImage(context context.Context, name string) error
	BuildImage(context context.Context, buildOptions types.ImageBuildOptions) error
	ListImages(context context.Context) ([]types.ImageSummary, error)
	InspectImage(context context.Context, name string) error
	PushImage(context context.Context, name string) error
}

//ContainerHandler  container handler
type ContainerHandler interface {
	CreateContainer(context context.Context, containerName string, config *containertypes.Config, hostConfig *containertypes.HostConfig, networkingConfig *networktypes.NetworkingConfig) error
	ContainerClean(context context.Context, container string) error
	StartContainer(context context.Context, container string) error
	StopContainer(context context.Context, container string) error
	RestartContainer(context context.Context, container string) error
	KillContainer(context context.Context, container string) error
	PauseContainer(context context.Context, container string) error
	GetContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
	//ListContainer()
	//InspectContainer()
	//ListContainerTopById(id string)
	//ListContainerStatsById(id string)
	//UpdateContainer()
	//ReNameContainer()
	//UnPauseContainer()
	//AttachContainer()
}

type Cli interface {
	ImageHandler
	ContainerHandler
	Close() error
}

const (
	caKey               = "ca.pem"
	certKey             = "cert.pem"
	keyKey              = "key.pem"
	pathSeparator       = string(os.PathSeparator)
	legacyDefaultDomain = "index.docker.io"
	defaultDomain       = "docker.io"
	officialRepoName    = "library"
)

//DockerCli
type DockerCli struct {
	client        dockerclient.APIClient
	in            *stream.InStream
	out           *stream.OutStream
	err           io.Writer
	IdentityToken string
}

var _ Cli = (*DockerCli)(nil)

//NewDefaultClient  new a default docker client
func NewDefaultDockerClient() (*DockerCli, error) {
	c, err := dockerclient.NewClientWithOpts(dockerclient.FromEnv)
	cli := &DockerCli{}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli.client = c
	cli.wrap()
	return cli, nil
}

// WithTLSClientConfig applies a tls config to the client transport.
func WithTLSClientConfig(cacertPath, certPath, keyPath string, InsecureSkipVerify bool) func(*dockerclient.Client) error {
	return func(c *dockerclient.Client) error {
		opts := tlsconfig.Options{
			CAFile:             cacertPath,
			CertFile:           certPath,
			KeyFile:            keyPath,
			ExclusiveRootPools: true,
			InsecureSkipVerify: InsecureSkipVerify,
		}
		config, err := tlsconfig.Client(opts)
		if err != nil {
			return errors.Wrap(err, "failed to create tls config")
		}
		if transport, ok := c.HTTPClient().Transport.(*http.Transport); ok {
			transport.TLSClientConfig = config
			return nil
		}
		return errors.Errorf("cannot apply tls config to transport: %T", c.HTTPClient().Transport)
	}
}

//NewClient new a customer  docker client
func NewDockerClient(dockerConfig *DockerConfig) (*DockerCli, error) {
	cli := &DockerCli{}
	host, err := opts.ValidateHost(dockerConfig.Host)
	if err != nil {
		return nil, err
	}
	clientOpts := []client.Opt{
		client.WithHost(host),
		client.WithVersion(""),
	}
	if dockerConfig.TLS {
		clientOpts = append(clientOpts, WithTLSClientConfig(
			fmt.Sprintf("%s%s%s", dockerConfig.CertDir, pathSeparator, caKey),
			fmt.Sprintf("%s%s%s", dockerConfig.CertDir, pathSeparator, certKey),
			fmt.Sprintf("%s%s%s", dockerConfig.CertDir, pathSeparator, keyKey),
			dockerConfig.InsecureSkipVerify,
		))
	}
	dockerClient, err := client.NewClientWithOpts(clientOpts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli.client = dockerClient
	cli.wrap()

	return cli, nil
}

//Wrap  reuse code by NewDockerClient
func (cli *DockerCli) wrap() {
	if cli.out == nil || cli.in == nil || cli.err == nil {
		stdin, stdout, stderr := term.StdStreams()
		if cli.in == nil {
			cli.in = stream.NewInStream(stdin)
		}
		if cli.out == nil {
			cli.out = stream.NewOutStream(stdout)
		}
		if cli.err == nil {
			cli.err = stderr
		}
	}
}

//Close close a docker client
func (cli *DockerCli) Close() error {
	if cli.client != nil {
		return errors.WithStack(cli.client.Close())
	}
	return nil
}

//PullImage pull a image  by image name
func (cli *DockerCli) PullImage(context context.Context, image string) error {
	domain, remainder := splitDockerDomain(image)
	imageName := domain + "/" + remainder

	responseBody, err := cli.client.ImagePull(context, imageName, types.ImagePullOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		err = responseBody.Close()
		if err != nil {
			logrus.Debugf("%+v", err)
		}
	}()
	return jsonmessage.DisplayJSONMessagesToStream(responseBody, cli.out, nil)
}

//RemoveImage remove a image by name
func (cli *DockerCli) RemoveImage(context context.Context, name string) error {
	err := cli.client.ContainerRemove(context, name, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//BuildImage build  a image
func (cli *DockerCli) BuildImage(context context.Context, buildOptions types.ImageBuildOptions) error {
	res, err := cli.client.ImageBuild(context, nil, buildOptions)
	if err != nil {
		return errors.WithStack(err)
	}
	defer res.Body.Close()
	return jsonmessage.DisplayJSONMessagesToStream(res.Body, cli.out, nil)
}

//ListImages  list docker  images
func (cli *DockerCli) ListImages(context context.Context) ([]types.ImageSummary, error) {

	images, err := cli.client.ImageList(context, types.ImageListOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return images, nil
}

//InspectImage see  image  inspect
func (cli *DockerCli) InspectImage(context context.Context, name string) error {
	getRefFunc := func(ref string) (interface{}, []byte, error) {
		return cli.client.ImageInspectWithRaw(context, ref)
	}
	err := inspect.Inspect(cli.out, []string{name}, "", getRefFunc)
	return errors.WithStack(err)
}

//RegistryLogin login
func (cli *DockerCli) RegistryLogin(context context.Context, auth types.AuthConfig) error {
	res, err := cli.client.RegistryLogin(context, auth)
	if err != nil {
		return errors.WithStack(err)
	}
	if res.Status == "200" {
		cli.IdentityToken = res.IdentityToken
	}
	return nil
}

//PushImage push a  image by  docker
func (cli *DockerCli) PushImage(context context.Context, image string) error {
	responseBody, err := cli.client.ImagePush(context, image, types.ImagePushOptions{All: true, RegistryAuth: cli.IdentityToken})
	if err != nil {
		return errors.WithStack(err)
	}

	return jsonmessage.DisplayJSONMessagesToStream(responseBody, cli.out, nil)
}

// splitDockerDomain splits a repository name to domain and remotename string.
// If no valid domain is found, the default domain is used. Repository name
// needs to be already validated before.
func splitDockerDomain(name string) (domain, remainder string) {
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ".:") && name[:i] != "localhost") {
		domain, remainder = defaultDomain, name
	} else {
		domain, remainder = name[:i], name[i+1:]
	}
	if domain == legacyDefaultDomain {
		domain = defaultDomain
	}
	if domain == defaultDomain && !strings.ContainsRune(remainder, '/') {
		remainder = officialRepoName + "/" + remainder
	}
	return
}

//CreateContainer create a Container by docker
func (cli *DockerCli) CreateContainer(context context.Context, containerName string, config *containertypes.Config, hostConfig *containertypes.HostConfig, networkingConfig *networktypes.NetworkingConfig) error {
	_, err := cli.client.ContainerCreate(context, config, hostConfig, networkingConfig, containerName)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//ContainerClean clean a ContainerClean
func (cli *DockerCli) ContainerClean(context context.Context, container string) error {
	if err := cli.client.ContainerRemove(context, container, types.ContainerRemoveOptions{Force: true}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//StartContainer start a container
func (cli *DockerCli) StartContainer(context context.Context, container string) error {
	err := cli.client.ContainerStart(context, container, types.ContainerStartOptions{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//StopContainer stop a container
func (cli *DockerCli) StopContainer(context context.Context, container string) error {
	var timeout *time.Duration
	timeoutValue := time.Duration(10) * time.Second
	timeout = &timeoutValue
	err := cli.client.ContainerStop(context, container, timeout)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//RestartContainer restart a container
func (cli *DockerCli) RestartContainer(context context.Context, container string) error {
	var timeout *time.Duration
	timeoutValue := time.Duration(10) * time.Second
	timeout = &timeoutValue
	err := cli.client.ContainerRestart(context, container, timeout)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//KillContainer  kill a container  by signal
func (cli *DockerCli) KillContainer(context context.Context, container string) error {
	if err := cli.client.ContainerKill(context, container, "kill"); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//PauseContainer  pause a container
func (cli *DockerCli) PauseContainer(context context.Context, container string) error {
	if err := cli.client.ContainerPause(context, container); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//GetContainerLogs  get a container logs
func (cli *DockerCli) GetContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	res, err := cli.client.ContainerLogs(ctx, container, options)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

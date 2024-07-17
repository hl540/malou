package container_runtime

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"io"
	"strings"
)

type DockerRuntime struct {
	*client.Client
}

func NewDockerRuntime(client *client.Client) (ContainerRuntime, error) {
	return &DockerRuntime{
		Client: client,
	}, nil
}

func (d *DockerRuntime) Create(ctx context.Context, imageName string, env []*EnvValue, workDir string) (string, error) {
	// 拉取镜像
	if err := d.pullImage(ctx, imageName); err != nil {
		return "", err
	}

	// 创建容器
	conf := &container.Config{
		Env:        EnvsToArray(env),
		Cmd:        []string{"sh"},
		Image:      imageName,
		WorkingDir: WorkDir,
		Tty:        true,
	}
	// 挂载目录
	hostConf := &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:%s", workDir, WorkDir)},
	}
	createResp, err := d.ContainerCreate(ctx, conf, hostConf, nil, nil, "")
	if err != nil {
		return "", err
	}
	// 启动容器
	if err := d.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {
		return "", err
	}
	return createResp.ID, nil
}

func (d *DockerRuntime) pullImage(ctx context.Context, imageName string) error {
	// 查询镜像是否在本地存在
	tag := strings.Split(imageName, ":")
	if len(tag) < 2 {
		tag = append(tag, "latest")
	}
	if tag[1] != "latest" {
		list, err := d.ImageList(ctx, image.ListOptions{Filters: filters.NewArgs(
			filters.Arg("reference", imageName),
		)})
		if err != nil {
			return err
		}
		if len(list) > 0 {
			return nil
		}
	}
	// 拉取镜像
	_, err := d.Client.ImagePull(ctx, imageName, image.PullOptions{})
	return err
}

func (d *DockerRuntime) AttachExec(ctx context.Context, containerID, cmd string) (io.Reader, error) {
	execResp, err := d.ContainerExecCreate(ctx, containerID, container.ExecOptions{
		AttachStderr: true,
		AttachStdout: true,
		WorkingDir:   WorkDir,
		Cmd:          []string{"sh", "-c", cmd},
	})
	if err != nil {
		return nil, err
	}
	attachResp, err := d.ContainerExecAttach(ctx, execResp.ID, container.ExecAttachOptions{})
	if err != nil {
		return nil, err
	}
	return attachResp.Reader, nil
}

func (d *DockerRuntime) Clear(ctx context.Context, containerID string) error {
	// 查询容器是否存在
	containers, err := d.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("id", containerID)),
	})
	if err != nil {
		return err
	}
	if len(containers) > 0 {
		return nil
	}
	// 删除容器，Force
	if err := d.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true}); err != nil {
		return err
	}
	return nil
}

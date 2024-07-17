package container_runtime

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	agent2 "github.com/hl540/malou/internal/app/agent"
	"strings"
)

type DockerRuntime struct {
	*client.Client
	out chan string
}

func NewDockerRuntime(client *client.Client) (ContainerRuntime, error) {
	return &DockerRuntime{
		Client: client,
		out:    make(chan string),
	}, nil
}

func (d *DockerRuntime) Create(ctx context.Context, imageName, workDir string) (string, error) {
	// 拉取镜像
	if err := d.pullImage(ctx, imageName); err != nil {
		return "", err
	}

	// 创建容器
	conf := &container.Config{
		Env:        agent2.GetEnvs(),
		Cmd:        []string{"sh"},
		Image:      imageName,
		WorkingDir: agent2.WorkDir,
		Tty:        true,
	}
	hostConf := &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:%s", workDir, agent2.WorkDir)},
	}
	createResp, err := d.ContainerCreate(ctx, conf, hostConf, nil, nil, "")
	if err != nil {
		return "", err
	}
	// 启动容器
	if err := d.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {
		if remErr := d.ContainerRemove(ctx, createResp.ID, container.RemoveOptions{}); remErr != nil {
			d.out <- fmt.Sprintf("Failed to remove container %s, %s", createResp.ID, remErr.Error())
		}
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
	out, err := d.Client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}

	type imagePullOut struct {
		Status string `json:"status"`
		ID     string `json:"id"`
	}
	// 获取日志输出
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		imagePullOutLine := imagePullOut{}
		if err := json.Unmarshal(scanner.Bytes(), &imagePullOutLine); err != nil {
			d.out <- fmt.Sprintf("Failed to pull image: %s", imageName)
		} else {
			d.out <- imagePullOutLine.Status
		}
	}
	return nil
}

func (d *DockerRuntime) AttachExec(ctx context.Context, containerID, cmd string) error {
	d.out <- fmt.Sprintf("Running %s", cmd)
	execResp, err := d.ContainerExecCreate(ctx, containerID, container.ExecOptions{
		AttachStderr: true,
		AttachStdout: true,
		Env:          agent2.GetEnvs(),
		WorkingDir:   agent2.WorkDir,
		Cmd:          []string{"sh", "-c", cmd},
	})
	if err != nil {
		return err
	}
	attachResp, err := d.ContainerExecAttach(ctx, execResp.ID, container.ExecAttachOptions{})
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(attachResp.Reader)
	for scanner.Scan() {
		d.out <- scanner.Text()
	}
	return nil
}

func (d *DockerRuntime) Clear(ctx context.Context, containerID string) error {
	// 停止容器
	timeout := 0
	stopOpt := container.StopOptions{Timeout: &timeout, Signal: "SIGKILL"}
	if err := d.ContainerStop(ctx, containerID, stopOpt); err != nil {
		return err
	}

	// 删除容器
	if err := d.ContainerRemove(ctx, containerID, container.RemoveOptions{}); err != nil {
		return err
	}
	close(d.out)
	return nil
}

func (d *DockerRuntime) OutLogCall() chan string {
	return d.out
}

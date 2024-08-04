package container_runtime

import (
	"context"
	"io"
)

const WorkDir = "/var/workspace"

type ContainerRuntime interface {
	Create(ctx context.Context, imageName string, env []*EnvValue, workDir string) (string, error)
	AttachExec(ctx context.Context, containerId, cmd string) (io.Reader, error)
	Clear(ctx context.Context, containerId string) error
}

type ContainerLog struct {
	ContainerId string
	Cmd         string
	Timestamp   int64
}

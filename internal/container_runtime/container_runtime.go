package container_runtime

import "context"

type ContainerRuntime interface {
	Create(ctx context.Context, imageName, workDir string) (string, error)
	AttachExec(ctx context.Context, containerID, cmd string) error
	Clear(ctx context.Context, containerID string) error
	OutLogCall(func(message string)) chan string
}

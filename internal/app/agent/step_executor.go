package agent

import (
	"context"
	"github.com/hl540/malou/internal/container_runtime"
	"github.com/hl540/malou/proto/v1"
	"log"
)

type StepExecutor interface {
	Execute(ctx context.Context, step *v1.Step, workDir string) error
}

type BaseStepExecutor struct {
	cr        container_runtime.ContainerRuntime
	reportLog *ReportLog
}

func NewBaseStepExecutor(cr container_runtime.ContainerRuntime, reportLog *ReportLog) StepExecutor {
	return &BaseStepExecutor{
		cr:        cr,
		reportLog: reportLog,
	}
}

func (e *BaseStepExecutor) Execute(ctx context.Context, step *v1.Step, workDir string) error {
	// 创建容器
	containerID, err := e.cr.Create(ctx, step.Image, workDir)
	if err != nil {
		return err
	}

	defer func() {
		err := e.cr.Clear(ctx, containerID)
		if err != nil {
			log.Printf("An error occurred to delete the %s: %s", containerID, err.Error())
		}
	}()

	// 多个命令Attach的方式依次执行
	for _, cmd := range step.Commands {
		if err := e.cr.AttachExec(ctx, containerID, cmd); err != nil {
			return err
		}
	}
	return nil
}

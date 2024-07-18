package runner

import (
	"bufio"
	"context"
	"github.com/hl540/malou/internal/runner/container_runtime"
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
	e.reportLog.WithStep(step.Name).Log("start executing step [%s]", step.Name)
	// 创建容器
	containerID, err := e.cr.Create(ctx, step.Image, nil, workDir)
	if err != nil {
		return err
	}
	defer func() {
		// 清理容器
		if err := e.cr.Clear(ctx, containerID); err != nil {
			log.Printf("An error occurred to delete the %s: %s", containerID, err.Error())
		}
	}()

	// 多个命令Attach的方式依次执行
	for _, cmd := range step.Commands {
		e.reportLog.WithStep(step.Name).WithCmd(cmd).Log("start executing cmd [%s]", cmd)
		out, err := e.cr.AttachExec(ctx, containerID, cmd)
		// 执行随便清理容器
		if err != nil {
			return err
		}
		// 逐行读取log
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			text := scanner.Text()
			e.reportLog.WithStep(step.Name).WithCmd(cmd).Log(text)
		}
	}
	return nil
}

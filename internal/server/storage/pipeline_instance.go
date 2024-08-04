package storage

import (
	"context"
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"time"
)

type PipelineInstanceModel struct {
	Id            string `db:"id"`
	PipelineId    int64  `db:"pipeline_id"`
	RuntimeConfig string `db:"runtime_config"`
	Status        int32  `db:"status"`
	StartTimeAt   int64  `db:"start_time_at"`
	Duration      int64  `db:"duration"`
	RunnerId      int64  `db:"runner_id"`
	CreatedAt     int64  `db:"created_at"`
	UpdatedAt     int64  `db:"updated_at"`
}

type PipelineInstanceEnvModel struct {
	Id                 int64  `db:"id"`
	PipelineInstanceId string `db:"pipeline_instance_id"`
	Name               string `db:"name"`
	Value              string `db:"value"`
}

type PipelineInstanceDao struct {
	Session
}

func NewPipelineInstanceDao(session Session) *PipelineInstanceDao {
	return &PipelineInstanceDao{Session: session}
}

// Create 创建pipeline实例
// transaction
func (d *PipelineInstanceDao) Create(ctx context.Context, pipelineInstance *PipelineInstanceModel, envs []*PipelineInstanceEnvModel) error {
	return TransactionCtx(ctx, func(ctx context.Context, tx Session) error {
		// 插入pipeline实例
		pipelineInstanceInsert := fmt.Sprintf("INSERT INTO %s (`id`, `pipeline_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)", PipelineInstanceTable)
		args := []any{
			pipelineInstance.Id,
			pipelineInstance.PipelineId,
			time.Now().Unix(),
			time.Now().Unix(),
		}
		_, err := d.ExecContext(ctx, pipelineInstanceInsert, args...)
		if err != nil {
			return err
		}

		// 插入pipeline实例env
		for _, env := range envs {
			env.PipelineInstanceId = pipelineInstance.Id
			pipelineInstanceEnvInsert := fmt.Sprintf("INSERT INTO %s (`pipeline_instance_id`, `name`, `value`) VALUES (?, ?, ?)", PipelineInstanceEnvTable)
			_, err := d.ExecContext(ctx, pipelineInstanceEnvInsert, env.PipelineInstanceId, env.Name, env.Value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Start 启动pipelineInstanceId，更新状态和运行时配置
func (d *PipelineInstanceDao) Start(ctx context.Context, pipelineInstanceId string, runtimeConfig string, runnerId int64) error {
	update := fmt.Sprintf("UPDATE %s SET `runtime_config` = ?, `status` = ?, `runner_id` = ? WHERE `id` = ? ", PipelineInstanceTable)
	args := []any{
		runtimeConfig,
		v1.PipelineInstanceType_Running,
		runnerId,
		pipelineInstanceId,
	}
	_, err := d.ExecContext(ctx, update, args...)
	if err != nil {
		return err
	}
	return nil
}

func (d *PipelineInstanceDao) GetPendingByRunnerId(ctx context.Context, runnerId int64) (*PipelineInstanceModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ? AND `status` = ?", PipelineInstanceTable)
	pipelineInstance := new(PipelineInstanceModel)
	err := d.GetContext(ctx, pipelineInstance, query, runnerId, v1.PipelineInstanceType_Pending)
	if err != nil {
		return nil, err
	}
	return pipelineInstance, nil
}

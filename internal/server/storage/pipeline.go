package storage

import (
	"context"
	"fmt"
	"time"
)

type PipelineModel struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type PipelineStepModel struct {
	ID         int64  `db:"id"`
	PipelineID int64  `db:"pipeline_id"`
	Name       string `db:"name"`
	Image      string `db:"image"`
	Commands   []string
}

type PipelineStepCmdModel struct {
	ID             int64  `db:"id"`
	PipelineID     int64  `db:"pipeline_id"`
	PipelineStepID int64  `db:"pipeline_step_id"`
	Cmd            string `db:"cmd"`
}

type PipelineDao struct {
	Session
}

func NewPipelineDao(session Session) *PipelineDao {
	return &PipelineDao{session}
}

// Create 创建pipeline
// transaction
func (d *PipelineDao) Create(ctx context.Context, pipeline *PipelineModel, steps []*PipelineStepModel) error {
	return TransactionCtx(ctx, func(ctx context.Context, tx Session) error {
		// 插入pipeline
		pipelineInsert := fmt.Sprintf("INSERT INTO %s (`name`, `created_at`, `updated_at`) VALUES (?, ?, ?)", PipelineTable)
		result, err := tx.ExecContext(ctx, pipelineInsert, pipeline.Name, time.Now().Unix(), time.Now().Unix())
		if err != nil {
			return err
		}
		pipeline.ID, err = result.LastInsertId()
		if err != nil {
			return err
		}
		// 插入step
		for _, step := range steps {
			step.PipelineID = pipeline.ID
			stepInsert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `name`, `image`) VALUES (?, ?, ?)", PipelineStepTable)
			result, err := tx.ExecContext(ctx, stepInsert, step.PipelineID, step.Name, step.Image)
			if err != nil {
				return err
			}
			step.ID, err = result.LastInsertId()
			if err != nil {
				return err
			}
			// 插入cmd
			for _, cmd := range step.Commands {
				cmdInsert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `pipeline_step_id`, `cmd`) VALUES (?, ?, ?)", PipelineStepCmdTable)
				if _, err := tx.ExecContext(ctx, cmdInsert, pipeline.ID, step.ID, cmd); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// Update 更新pipeline
// transaction
func (d *PipelineDao) Update(ctx context.Context, pipeline *PipelineModel, steps []*PipelineStepModel) error {
	return TransactionCtx(ctx, func(ctx context.Context, tx Session) error {
		// 更新pipeline
		pipelineUpdate := fmt.Sprintf("UPDATE %s SET `name` = ?, `updated_at` = ? WHERE id= ?", PipelineTable)
		_, err := tx.ExecContext(ctx, pipelineUpdate, pipeline.Name, time.Now().Unix(), pipeline.ID)
		if err != nil {
			return err
		}
		// 清理原来的step和cmd
		delQuery := fmt.Sprintf("DELETE FROM %s where pipeline_id = ?", PipelineStepCmdTable)
		if _, err := tx.ExecContext(ctx, delQuery, pipeline.ID); err != nil {
			return err
		}
		delQuery = fmt.Sprintf("DELETE FROM %s where pipeline_id = ?", PipelineStepTable)
		if _, err := tx.ExecContext(ctx, delQuery, pipeline.ID); err != nil {
			return err
		}
		// 插入step
		for _, step := range steps {
			step.PipelineID = pipeline.ID
			stepInsert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `name`, `image`) VALUES (?, ?, ?)", PipelineStepTable)
			result, err := tx.ExecContext(ctx, stepInsert, step.PipelineID, step.Name, step.Image)
			if err != nil {
				return err
			}
			step.ID, err = result.LastInsertId()
			if err != nil {
				return err
			}
			// 插入cmd
			for _, cmd := range step.Commands {
				cmdInsert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `pipeline_step_id`, `cmd`) VALUES (?, ?, ?)", PipelineStepCmdTable)
				if _, err := tx.ExecContext(ctx, cmdInsert, pipeline.ID, step.ID, cmd); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (d *PipelineDao) GetByID(ctx context.Context, id int64) (*PipelineModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", PipelineTable)
	var pipeline PipelineModel
	err := d.GetContext(ctx, &pipeline, query, id)
	if err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (d *PipelineDao) GetInfoByID(ctx context.Context, id int64) (*PipelineModel, []*PipelineStepModel, error) {
	pipelineQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", PipelineTable)
	pipeline := new(PipelineModel)
	err := d.GetContext(ctx, pipeline, pipelineQuery, id)
	if err != nil {
		return nil, nil, err
	}

	var steps []*PipelineStepModel
	stepsQuery := fmt.Sprintf("SELECT * FROM %s WHERE pipeline_id = ?", PipelineStepTable)
	err = d.SelectContext(ctx, &steps, stepsQuery, id)
	if err != nil {
		return nil, nil, err
	}
	for _, step := range steps {
		cmdQuery := fmt.Sprintf("SELECT * FROM %s WHERE pipeline_step_id = ?", PipelineStepCmdTable)
		var commands []*PipelineStepCmdModel
		err = d.SelectContext(ctx, &commands, cmdQuery, step.ID)
		if err != nil {
			return nil, nil, err
		}
		for _, cmd := range commands {
			step.Commands = append(step.Commands, cmd.Cmd)
		}
	}
	return pipeline, steps, nil
}

type PipelineSearchListParam struct {
	Name string
	Page int64
	Size int64
}

func (d *PipelineDao) SearchList(ctx context.Context, param *PipelineSearchListParam) ([]*PipelineModel, int64, error) {
	query := fmt.Sprintf("FROM %s WHERE 1 = 1", PipelineTable)
	args := make([]any, 0)
	if param.Name != "" {
		query = fmt.Sprintf("%s AND `name` LIKE ?", query)
		args = append(args, `%`+param.Name+`%`)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(`id`) %s", query)
	var count int64
	if err := d.GetContext(ctx, &count, countQuery, args...); err != nil {
		return nil, 0, err
	}
	listQuery := fmt.Sprintf("SELECT * %s LIMIT %d, %d", query, (param.Page-1)*param.Size, param.Size)

	var pipelines []*PipelineModel
	err := d.SelectContext(ctx, &pipelines, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	return pipelines, count, nil
}

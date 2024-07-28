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

func (d *PipelineDao) GetByID(ctx context.Context, id int64) (*PipelineModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", PipelineTable)
	var pipeline PipelineModel
	err := d.GetContext(ctx, &pipeline, query, id)
	if err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (d *PipelineDao) Create(ctx context.Context, data *PipelineModel) error {
	// 插入pipeline
	insert := fmt.Sprintf("INSERT INTO %s (`name`, `created_at`, `updated_at`) VALUES (?)", PipelineTable)
	args := []any{data.Name, time.Now().Unix(), time.Now().Unix()}
	result, err := d.ExecContext(ctx, insert, args...)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}

func (d *PipelineDao) Update(ctx context.Context, data *PipelineModel) error {
	// 插入pipeline
	update := fmt.Sprintf("UPDATE %s SET `name` = ?, `updated_at` = ? WHERE id= ?", PipelineTable)
	args := []any{data.Name, time.Now().Unix(), data.ID}
	_, err := d.ExecContext(ctx, update, args...)
	if err != nil {
		return err
	}
	// 清理原来的step和cmd
	delQuery := fmt.Sprintf("DELETE FROM %s where pipeline_id = ?", PipelineStepCmdTable)
	if _, err := d.ExecContext(ctx, delQuery, data.ID); err != nil {
		return err
	}
	delQuery = fmt.Sprintf("DELETE FROM %s where pipeline_id = ?", PipelineStepTable)
	if _, err := d.ExecContext(ctx, delQuery, data.ID); err != nil {
		return err
	}
	return nil
}

func (d *PipelineDao) BatchSavePipelineStep(ctx context.Context, pipelineID int64, data []*PipelineStepModel) error {
	for _, step := range data {
		step.PipelineID = pipelineID
		if err := d.SavePipelineStep(ctx, step); err != nil {
			return err
		}
	}
	return nil
}

func (d *PipelineDao) SavePipelineStep(ctx context.Context, data *PipelineStepModel) error {
	// 插入step
	insert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `name`, `image`) VALUES (?, ?, ?)", PipelineStepTable)
	result, err := d.ExecContext(ctx, insert, data.PipelineID, data.Name, data.Image)
	if err != nil {
		return err
	}
	stepID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	data.ID = stepID

	// 插入step执行的命令
	for _, cmd := range data.Commands {
		insert := fmt.Sprintf("INSERT INTO %s (`pipeline_id`, `pipeline_step_id`, `cmd`) VALUES (?, ?, ?)", PipelineStepCmdTable)
		if _, err := d.ExecContext(ctx, insert, data.PipelineID, stepID, cmd); err != nil {
			return err
		}
	}
	return nil
}

func (d *PipelineDao) GetStepsByPipelineId(ctx context.Context, pipelineID int64) ([]*PipelineStepModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE pipeline_id = ?", PipelineStepTable)
	var steps []*PipelineStepModel
	err := d.SelectContext(ctx, &steps, query, pipelineID)
	if err != nil {
		return nil, err
	}
	for _, step := range steps {
		query := fmt.Sprintf("SELECT * FROM %s WHERE pipeline_step_id = ?", PipelineStepCmdTable)
		var commands []*PipelineStepCmdModel
		err := d.SelectContext(ctx, &commands, query, step.ID)
		if err != nil {
			return nil, err
		}
		for _, cmd := range commands {
			step.Commands = append(step.Commands, cmd.Cmd)
		}
	}
	return steps, nil
}

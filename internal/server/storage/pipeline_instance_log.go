package storage

import (
	"context"
	"fmt"
)

type PipelineInstanceLogModel struct {
	PipelineInstanceID string `db:"pipeline_instance_id"`
	StepName           string `db:"step_name"`
	Cmd                string `db:"cmd"`
	Result             string `db:"result"`
	Type               string `db:"type"`
	Timestamp          int64  `db:"timestamp"`
	Duration           int64  `db:"duration"`
}

type PipelineInstanceLogDao struct {
	Session
}

func NewPipelineInstanceLogDao(session Session) *PipelineInstanceLogDao {
	return &PipelineInstanceLogDao{Session: session}
}

func (d *PipelineInstanceLogDao) Insert(ctx context.Context, data *PipelineInstanceLogModel) error {
	insert := fmt.Sprintf("INSERT INTO %s (`pipeline_instance_id`, `step_name`, `cmd`, `result`, `type`, `timestamp`, `duration`) VALUES (?, ?, ?, ?, ?, ?, ?)", PipelineInstanceLogTable)
	args := []any{
		data.PipelineInstanceID,
		data.StepName,
		data.Cmd,
		data.Result,
		data.Type,
		data.Timestamp,
		data.Duration,
	}
	_, err := d.ExecContext(ctx, insert, args...)
	return err
}

func (d *PipelineInstanceLogDao) DelByPipelineInstanceID(ctx context.Context, pipelineInstanceID string) error {
	delQuery := fmt.Sprintf("DELETE FROM %s WHERE `pipeline_instance_id` = ?", PipelineInstanceLogTable)
	_, err := d.ExecContext(ctx, delQuery, pipelineInstanceID)
	return err
}

func (d *PipelineInstanceLogDao) GetLogsByPipelineInstanceID(ctx context.Context, pipelineInstanceID string, offset int64) ([]*PipelineInstanceLogModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `pipeline_instance_id` = ? OFFSET ?", PipelineInstanceLogTable)
	var logs []*PipelineInstanceLogModel
	err := d.SelectContext(ctx, &logs, query, pipelineInstanceID, offset)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

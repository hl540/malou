package storage

import (
	"context"
	"fmt"
)

type PipelineInstanceLogModel struct {
	PipelineInstanceId string `db:"pipeline_instance_id"`
	Serial             int64  `db:"serial"`
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
	insert := fmt.Sprintf("INSERT INTO %s (`pipeline_instance_id`, `serial`, `step_name`, `cmd`, `result`, `type`, `timestamp`, `duration`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", PipelineInstanceLogTable)
	args := []any{
		data.PipelineInstanceId,
		data.Serial,
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

func (d *PipelineInstanceLogDao) GetLogsByPipelineInstanceId(ctx context.Context, pipelineInstanceId string, offset int64) ([]*PipelineInstanceLogModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `pipeline_instance_id` = ? OFFSET ?", PipelineInstanceLogTable)
	var logs []*PipelineInstanceLogModel
	err := d.SelectContext(ctx, &logs, query, pipelineInstanceId, offset)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

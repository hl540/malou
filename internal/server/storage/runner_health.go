package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type RunnerHealthModel struct {
	RunnerID          int64   `db:"runner_id"`
	CpuPercent        float64 `db:"cpu_percent"`
	MemoryTotal       float64 `db:"memory_total"`
	MemoryUsed        float64 `db:"memory_used"`
	MemoryFree        float64 `db:"memory_free"`
	MemoryUsedPercent float64 `db:"memory_used_percent"`
	DiskTotal         float64 `db:"disk_total"`
	DiskUsed          float64 `db:"disk_used"`
	DiskFree          float64 `db:"disk_free"`
	DiskUsedPercent   float64 `db:"disk_used_percent"`
	WorkerStatus      string  `db:"worker_status"`
	CreatedAt         int64   `db:"created_at"`
}

type RunnerHealthDao struct {
	Session
}

func NewRunnerHealthDao(session Session) *RunnerHealthDao {
	return &RunnerHealthDao{session}
}

func (d *RunnerHealthDao) Insert(ctx context.Context, data *RunnerHealthModel) error {
	err := d.cleanOverdue(ctx, data.RunnerID)
	if err != nil {
		logrus.WithContext(ctx).Warningf("cleanOverdue: %v", err)
	}

	insert := fmt.Sprintf("INSERT INTO %s (`runner_id`, `cpu_percent`, `memory_total`, `memory_used`, `memory_free`, `memory_used_percent`, `disk_total`, `disk_used`, `disk_free`, `disk_used_percent`, `worker_status`, `created_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", RunnerHealthTable)
	_, err = d.ExecContext(ctx, insert, []any{
		data.RunnerID,
		data.CpuPercent,
		data.MemoryTotal,
		data.MemoryUsed,
		data.MemoryFree,
		data.MemoryUsedPercent,
		data.DiskTotal,
		data.DiskUsed,
		data.DiskFree,
		data.DiskUsedPercent,
		data.WorkerStatus,
		data.CreatedAt,
	}...)
	return err
}

func (d *RunnerHealthDao) cleanOverdue(ctx context.Context, runnerID int64) error {
	query := fmt.Sprintf("SELECT `created_at` FROM %s WHERE runner_id = ? ORDER BY created_at DESC LIMIT 1 OFFSET 9", RunnerHealthTable)
	var createdAt int64
	err := d.GetContext(ctx, &createdAt, query, runnerID)
	if err != nil {
		return err
	}
	delQuery := fmt.Sprintf("DELETE FROM %s WHERE `runner_id` = ? AND `created_at` <= ?", RunnerHealthTable)
	_, err = d.ExecContext(ctx, delQuery, runnerID, createdAt)
	if err != nil {
		return err
	}
	return err
}

func (d *RunnerHealthDao) GetAllByRunnerID(ctx context.Context, runnerID int64) ([]*RunnerHealthModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE runner_id = ?", RunnerHealthTable)
	var healths []*RunnerHealthModel
	err := d.SelectContext(ctx, &healths, query, runnerID)
	if err != nil {
		return nil, err
	}
	return healths, nil
}

func (d *RunnerHealthDao) GetLatestByRunnerID(ctx context.Context, runnerID int64) (*RunnerHealthModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE runner_id = ? ORDER BY `runner_id` LIMIT 1", RunnerHealthTable)
	var health RunnerHealthModel
	err := d.GetContext(ctx, &health, query, runnerID)
	if err != nil {
		return nil, err
	}
	return &health, nil
}

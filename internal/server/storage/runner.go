package storage

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type RunnerModel struct {
	ID        int64  `db:"id"`
	Code      string `db:"code"`
	Secret    string `db:"secret"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type RunnerLabelModel struct {
	ID         int64  `db:"id"`
	RunnerCode string `db:"runner_code"`
	Value      string `db:"value"`
}

type RunnerDao struct {
	Session
}

func NewRunnerDao(session Session) *RunnerDao {
	return &RunnerDao{Session: session}
}

func (d *RunnerDao) Add(ctx context.Context, runner *RunnerModel) error {
	query := fmt.Sprintf("INSERT INTO %s (`code`, `secret`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", RunnerTable)
	args := []any{runner.Code, runner.Secret, runner.Name, time.Now().Unix(), time.Now().Unix()}
	result, err := d.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	runner.ID, err = result.LastInsertId()
	return err
}

func (d *RunnerDao) GetByCode(ctx context.Context, code string) (*RunnerModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE code = ? LIMIT 1", RunnerTable)
	runner := new(RunnerModel)
	err := d.GetContext(ctx, runner, query, code)
	return runner, err
}

func (d *RunnerDao) SaveLabel(ctx context.Context, runnerCode string, labels []string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE runner_code = ?", RunnerLabelTable)
	_, err := d.ExecContext(ctx, query, runnerCode)
	if err != nil {
		return err
	}
	if len(labels) == 0 {
		return nil
	}
	query = fmt.Sprintf("INSERT INTO %s (runner_code, labels) VALUES (?, ?)", RunnerTable)
	var args []any
	var temps []string
	for _, label := range labels {
		args = append(args, runnerCode, label)
		temps = append(temps, "(?, ?)")
	}
	query += strings.Join(temps, ", ")
	_, err = d.ExecContext(ctx, query, args...)
	return err
}

package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
)

type RunnerModel struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type RunnerDao struct {
	Session
	table string
}

func NewRunnerDao(session Session) *RunnerDao {
	return &RunnerDao{
		Session: session,
		table:   "ml_runner",
	}
}

func (d *RunnerDao) Add(ctx context.Context, runner *RunnerModel) (string, error) {
	sql, args, _ := squirrel.Insert(d.table).SetMap(map[string]interface{}{
		"id":   runner.ID,
		"name": runner.Name,
	}).ToSql()
	if _, err := d.ExecContext(ctx, sql, args); err != nil {
		return "", err
	}
	return runner.ID, nil
}

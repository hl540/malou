package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type RunnerModel struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type RunnerDao struct {
	*Base
	table string
}

func NewRunnerDao(db *sqlx.DB) *RunnerDao {
	return &RunnerDao{
		Base:  NewBase(db),
		table: "ml_runner",
	}
}

func (d *RunnerDao) Add(ctx context.Context, runner *RunnerModel) (string, error) {
	_, err := squirrel.Insert(d.table).Values(
		runner.ID,
		runner.Name,
	).RunWith(d.DB).Exec()
	if err != nil {
		return "", err
	}
	return runner.ID, nil
}

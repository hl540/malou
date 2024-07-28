package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type RunnerModel struct {
	ID        int64  `db:"id"`
	Code      string `db:"code"`
	Secret    string `db:"secret"`
	Name      string `db:"name"`
	Status    int32  `db:"status"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type RunnerLabelModel struct {
	ID       int64  `db:"id"`
	RunnerID int64  `db:"runner_id"`
	Value    string `db:"value"`
}

type RunnerEnvModel struct {
	ID       int64  `db:"id"`
	RunnerID int64  `db:"runner_id"`
	Name     string `db:"name"`
	Value    string `db:"value"`
}

type RunnerDao struct {
	Session
}

func NewRunnerDao(session Session) *RunnerDao {
	return &RunnerDao{Session: session}
}

func (d *RunnerDao) Create(ctx context.Context, runner *RunnerModel) error {
	query := fmt.Sprintf("INSERT INTO %s (`code`, `secret`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", RunnerTable)
	args := []any{runner.Code, runner.Secret, runner.Name, time.Now().Unix(), time.Now().Unix()}
	result, err := d.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	runner.ID, err = result.LastInsertId()
	return err
}

func (d *RunnerDao) Update(ctx context.Context, runner *RunnerModel) error {
	update := fmt.Sprintf("UPDATE %s SET `name` = ?, `updated_at` = ? WHERE `id` = ?", RunnerTable)
	args := []any{runner.Name, time.Now().Unix(), runner.ID}
	_, err := d.ExecContext(ctx, update, args...)
	if err != nil {
		return err
	}
	// 清理原来的label和env
	delQuery := fmt.Sprintf("DELETE FROM %s WHERE `runner_id` = ?", RunnerLabelTable)
	if _, err := d.ExecContext(ctx, delQuery, runner.ID); err != nil {
		return err
	}
	delQuery = fmt.Sprintf("DELETE FROM %s WHERE `runner_id` = ?", RunnerEnvTable)
	if _, err := d.ExecContext(ctx, delQuery, runner.ID); err != nil {
		return err
	}
	return nil
}

func (d *RunnerDao) GetByID(ctx context.Context, id int64) (*RunnerModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `id` = ? LIMIT 1", RunnerTable)
	runner := new(RunnerModel)
	err := d.GetContext(ctx, runner, query, id)
	return runner, err
}

func (d *RunnerDao) GetByCode(ctx context.Context, code string) (*RunnerModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `code` = ? LIMIT 1", RunnerTable)
	runner := new(RunnerModel)
	err := d.GetContext(ctx, runner, query, code)
	return runner, err
}

func (d *RunnerDao) GetList(ctx context.Context, code, name string, labels []string, page, size int64) ([]*RunnerModel, int64, error) {
	query := fmt.Sprintf("FROM %s AS r LEFT JOIN %s AS rl ON r.`id` = rl.runner_id WHERE 1 = 1", RunnerTable, RunnerLabelTable)
	args := make([]any, 0)
	if code != "" {
		query = fmt.Sprintf("%s AND r.`code` LIKE ?", query)
		args = append(args, `%`+code+`%`)
	}
	if name != "" {
		query = fmt.Sprintf("%s AND r.`name` LIKE ?", query)
		args = append(args, `%`+name+`%`)
	}
	if len(labels) > 0 {
		inQuery, inArgs, _ := sqlx.In("?", labels)
		query = fmt.Sprintf("%s AND rl.`value` IN (%s)", query, inQuery)
		args = append(args, inArgs...)
	}
	countQuery := fmt.Sprintf("SELECT COUNT(DISTINCT r.`code`) %s", query)
	var count int64
	if err := d.GetContext(ctx, &count, countQuery, args...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("SELECT DISTINCT r.* %s LIMIT %d, %d", query, (page-1)*size, size)
	var runners []*RunnerModel
	err := d.SelectContext(ctx, &runners, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	return runners, count, nil
}

func (d *RunnerDao) GetLabelByRunnerID(ctx context.Context, runnerID string) ([]*RunnerLabelModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ?", RunnerLabelTable)
	var labels []*RunnerLabelModel
	err := d.SelectContext(ctx, &labels, query, runnerID)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

// SaveLabel 保存标签
func (d *RunnerDao) SaveLabel(ctx context.Context, runnerID int64, labels []string) error {
	if len(labels) == 0 {
		return nil
	}
	query := fmt.Sprintf("INSERT INTO %s (`runner_id`, `value`) VALUES ", RunnerLabelTable)
	var args []any
	var temps []string
	for _, label := range labels {
		args = append(args, runnerID, label)
		temps = append(temps, "(?, ?)")
	}
	query += strings.Join(temps, ", ")
	_, err := d.ExecContext(ctx, query, args...)
	return err
}

func (d *RunnerDao) GetEnvByRunnerID(ctx context.Context, runnerID int64) ([]*RunnerEnvModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ?", RunnerEnvTable)
	var envs []*RunnerEnvModel
	err := d.SelectContext(ctx, &envs, query, runnerID)
	if err != nil {
		return nil, err
	}
	return envs, nil
}

// SaveEnv 保存环境变量
func (d *RunnerDao) SaveEnv(ctx context.Context, runnerID int64, env map[string]string) error {
	if len(env) == 0 {
		return nil
	}
	query := fmt.Sprintf("INSERT INTO %s (`runner_id`, `name`, `value`) VALUES ", RunnerEnvTable)
	var args []any
	var temps []string
	for name, value := range env {
		args = append(args, runnerID, name, value)
		temps = append(temps, "(?, ?, ?)")
	}
	query += strings.Join(temps, ", ")
	_, err := d.ExecContext(ctx, query, args...)
	return err
}

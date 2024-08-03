package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type RunnerModel struct {
	Id        int64  `db:"id"`
	Code      string `db:"code"`
	Secret    string `db:"secret"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type RunnerLabelModel struct {
	Id       int64  `db:"id"`
	RunnerId int64  `db:"runner_id"`
	Value    string `db:"value"`
}

type RunnerEnvModel struct {
	Id       int64  `db:"id"`
	RunnerId int64  `db:"runner_id"`
	Name     string `db:"name"`
	Value    string `db:"value"`
}

type RunnerDao struct {
	Session
}

func NewRunnerDao(session Session) *RunnerDao {
	return &RunnerDao{Session: session}
}

// Create 创建runner
// transaction
func (d *RunnerDao) Create(ctx context.Context, runner *RunnerModel, labels []*RunnerLabelModel, envs []*RunnerEnvModel) (err error) {
	return TransactionCtx(ctx, func(ctx context.Context, tx Session) error {
		runner.CreatedAt = time.Now().Unix()
		runner.UpdatedAt = time.Now().Unix()
		runnerInsert := fmt.Sprintf("INSERT INTO %s (`code`, `secret`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", RunnerTable)
		runnerInsertArgs := []any{runner.Code, runner.Secret, runner.Name, time.Now().Unix(), time.Now().Unix()}
		result, err := tx.ExecContext(ctx, runnerInsert, runnerInsertArgs...)
		if err != nil {
			return err
		}
		runner.Id, err = result.LastInsertId()
		if err != nil {
			return err
		}
		// 插入label
		for _, label := range labels {
			label.RunnerId = runner.Id
			labelInsert := fmt.Sprintf("INSERT INTO %s (`runner_id`, `value`) VALUES (?, ?)", RunnerLabelTable)
			_, err := tx.ExecContext(ctx, labelInsert, label.RunnerId, label.Value)
			if err != nil {
				return err
			}
		}
		// 插入env
		for _, env := range envs {
			env.RunnerId = runner.Id
			envInsert := fmt.Sprintf("INSERT INTO %s (`runner_id`, `name`, `value`) VALUES (?, ?, ?)", RunnerEnvTable)
			_, err := tx.ExecContext(ctx, envInsert, env.RunnerId, env.Name, env.Value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Update 更新runner
// transaction
func (d *RunnerDao) Update(ctx context.Context, runner *RunnerModel, labels []*RunnerLabelModel, envs []*RunnerEnvModel) error {
	return TransactionCtx(ctx, func(ctx context.Context, tx Session) error {
		update := fmt.Sprintf("UPDATE %s SET `name` = ?, `updated_at` = ? WHERE `id` = ?", RunnerTable)
		args := []any{runner.Name, time.Now().Unix(), runner.Id}
		_, err := tx.ExecContext(ctx, update, args...)
		if err != nil {
			return err
		}
		// 清理原来的label和env
		delQuery := fmt.Sprintf("DELETE FROM %s WHERE `runner_id` = ?", RunnerLabelTable)
		if _, err := tx.ExecContext(ctx, delQuery, runner.Id); err != nil {
			return err
		}
		delQuery = fmt.Sprintf("DELETE FROM %s WHERE `runner_id` = ?", RunnerEnvTable)
		if _, err := tx.ExecContext(ctx, delQuery, runner.Id); err != nil {
			return err
		}
		// 插入label
		for _, label := range labels {
			label.RunnerId = runner.Id
			labelInsert := fmt.Sprintf("INSERT INTO %s (`runner_id`, `value`) VALUES (?, ?)", RunnerLabelTable)
			_, err := tx.ExecContext(ctx, labelInsert, label.RunnerId, label.Value)
			if err != nil {
				return err
			}
		}
		// 插入env
		for _, env := range envs {
			env.RunnerId = runner.Id
			envInsert := fmt.Sprintf("INSERT INTO %s (`runner_id`, `name`, `value`) VALUES (?, ?, ?)", RunnerEnvTable)
			_, err := tx.ExecContext(ctx, envInsert, env.RunnerId, env.Name, env.Value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

type RunnerSearchListParam struct {
	Code   string
	Name   string
	Labels []string
	Page   int64
	Size   int64
}

func (d *RunnerDao) SearchList(ctx context.Context, param *RunnerSearchListParam) ([]*RunnerModel, int64, error) {
	query := fmt.Sprintf("FROM %s AS r LEFT JOIN %s AS rl ON r.`id` = rl.runner_id WHERE 1 = 1", RunnerTable, RunnerLabelTable)
	args := make([]any, 0)
	if param.Code != "" {
		query = fmt.Sprintf("%s AND r.`code` LIKE ?", query)
		args = append(args, `%`+param.Code+`%`)
	}
	if param.Name != "" {
		query = fmt.Sprintf("%s AND r.`name` LIKE ?", query)
		args = append(args, `%`+param.Name+`%`)
	}
	if len(param.Labels) > 0 {
		inQuery, inArgs, _ := sqlx.In("?", param.Labels)
		query = fmt.Sprintf("%s AND rl.`value` IN (%s)", query, inQuery)
		args = append(args, inArgs...)
	}
	countQuery := fmt.Sprintf("SELECT COUNT(DISTINCT r.`code`) %s", query)
	var count int64
	if err := d.GetContext(ctx, &count, countQuery, args...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("SELECT DISTINCT r.* %s LIMIT %d, %d", query, (param.Page-1)*param.Size, param.Size)
	var runners []*RunnerModel
	err := d.SelectContext(ctx, &runners, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	return runners, count, nil
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

func (d *RunnerDao) GetInfoById(ctx context.Context, id int64) (*RunnerModel, []*RunnerLabelModel, []*RunnerEnvModel, error) {
	queryRunner := fmt.Sprintf("SELECT * FROM %s WHERE `id` = ? LIMIT 1", RunnerTable)
	runner := new(RunnerModel)
	err := d.GetContext(ctx, runner, queryRunner, id)
	if err != nil {
		return nil, nil, nil, err
	}

	queryLabels := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ?", RunnerLabelTable)
	var labels []*RunnerLabelModel
	err = d.SelectContext(ctx, &labels, queryLabels, id)
	if err != nil {
		return nil, nil, nil, err
	}

	queryEnvs := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ?", RunnerEnvTable)
	var envs []*RunnerEnvModel
	err = d.SelectContext(ctx, &envs, queryEnvs, id)
	if err != nil {
		return nil, nil, nil, err
	}
	return runner, labels, envs, nil
}

func (d *RunnerDao) GetLabelsByRunnerId(ctx context.Context, runnerId int64) ([]string, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE `runner_id` = ?", RunnerLabelTable)
	query, args, err := sqlx.In(query, runnerId)
	if err != nil {
		return nil, err
	}
	var labels []*RunnerLabelModel
	err = d.SelectContext(ctx, &labels, query, args...)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(labels))
	for _, label := range labels {
		result = append(result, label.Value)
	}
	return result, nil
}

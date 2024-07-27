package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hl540/malou/internal/server"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

var (
	Runner *RunnerDao
)

const (
	RunnerTable      = "ml_runner"
	RunnerLabelTable = "ml_runner_label"
)

func InitDB(config *server.Config) (*sqlx.DB, error) {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/malou")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	Runner = NewRunnerDao(db)
	return db, nil
}

type Session interface {
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Execer
	sqlx.ExecerContext
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func TransactionCtx(ctx context.Context, fn func(ctx context.Context, tx Session) error) (err error) {
	var tx *sqlx.Tx
	tx, err = db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("recover from %#v, rollback failed: %w", p, e)
			} else {
				err = fmt.Errorf("recover from %#v", p)
			}
		} else if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("transaction failed: %s, rollback failed: %w", err, e)
			}
		} else {
			err = tx.Commit()
		}
	}()
	return fn(ctx, tx)
}

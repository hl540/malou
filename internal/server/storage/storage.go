package storage

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hl540/malou/internal/server"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

var (
	Runner *RunnerDao
)

func InitDB(config *server.Config) error {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/malou")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	Runner = NewRunnerDao(db)
	return nil
}

type Base struct {
	*sqlx.DB
}

func NewBase(db *sqlx.DB) *Base {
	return &Base{
		DB: db,
	}
}

func (b *Base) TransactCtx(ctx context.Context, fn func(ctx context.Context, tx *sqlx.Tx) error) (err error) {
	var tx *sqlx.Tx
	tx, err = b.BeginTxx(ctx, nil)
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

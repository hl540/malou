package storage

import (
	"context"
	"github.com/hl540/malou/internal/server"
	"testing"
)

func TestAddRunner(t *testing.T) {
	InitDB(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})

	t.Run("TestAddRunner", func(t *testing.T) {
		got, err := Runner.Add(context.Background(), &RunnerModel{
			ID:   "xxxxx",
			Name: "sssss",
		})
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(got)
	})

	t.Run("TestAddRunner", func(t *testing.T) {
		err := TransactCtx(context.Background(), func(ctx context.Context, tx Session) error {
			dao := NewRunnerDao(tx)
			_, err := dao.Add(ctx, &RunnerModel{
				ID:   "xxxxx",
				Name: "sssss",
			})
			if err != nil {
				return err
			}
			_, err = dao.Add(ctx, &RunnerModel{
				ID:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Name: "sssss",
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
	})
}

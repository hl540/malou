package test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/internal/server/storage"
	"github.com/hl540/malou/utils"
	"testing"
)

func init() {
	_, err := storage.InitDB(&server.Config{
		DBDrive:  "mysql",
		DBSource: "root:123456@tcp(127.0.0.1:3306)/malou",
	})
	if err != nil {
		panic(err)
	}
}

func TestAddRunner(t *testing.T) {
	t.Run("TestAddRunner", func(t *testing.T) {
		data := &storage.RunnerModel{
			Code:   uuid.New().String(),
			Secret: utils.StringWithCharsetV4(20),
			Name:   "runner_1",
		}
		err := storage.Runner.Create(context.Background(), data)
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(data)
	})
}

func TestRunnerDao_GetByCode(t *testing.T) {
	t.Run("GetByCode", func(t *testing.T) {
		got, err := storage.Runner.GetByCode(context.Background(), "61729cc5-12a4-4363-be5a-89e5e3a85dbb")
		fmt.Printf("%v", errors.Is(err, sql.ErrNoRows))
		if err != nil {
			t.Errorf("GetByCode() error = %v", err)
			return
		}
		t.Log(got)
	})

}

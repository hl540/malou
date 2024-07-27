package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/utils"
	"testing"
)

func TestAddRunner(t *testing.T) {
	InitDB(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})

	t.Run("TestAddRunner", func(t *testing.T) {
		data := &RunnerModel{
			Code:   uuid.New().String(),
			Secret: utils.StringWithCharsetV4(20),
			Name:   "runner_1",
		}
		err := Runner.Add(context.Background(), data)
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(data)
	})
}

func TestRunnerDao_GetByCode(t *testing.T) {
	InitDB(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})
	t.Run("GetByCode", func(t *testing.T) {
		got, err := Runner.GetByCode(context.Background(), "61729cc5-12a4-4363-be5a-89e5e3a85dbb")
		fmt.Printf("%v", errors.Is(err, sql.ErrNoRows))
		if err != nil {
			t.Errorf("GetByCode() error = %v", err)
			return
		}
		t.Log(got)
	})

}

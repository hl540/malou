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
}

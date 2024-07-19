package storage

import (
	"context"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/utils"
	"testing"
)

func TestGetRunnerList(t *testing.T) {
	InitMongo(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})
	t.Run("TestGetRunnerList", func(t *testing.T) {
		got, num, err := GetRunnerList(context.Background(), &RunnerListOption{
			Name:   "7",
			Labels: []string{"l5", "l1", "l9"},
		})
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(num)
		t.Log(got)
	})
}

func TestGetRunnerByID(t *testing.T) {
	InitMongo(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})
	t.Run("TestGetRunnerByID", func(t *testing.T) {
		got, err := GetRunnerByID(context.Background(), "669a0be70be1dea32aeeb04a")
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(got)
	})
}

func TestAddRunner(t *testing.T) {
	InitMongo(&server.Config{
		MongoUri:      "mongodb://localhost:27017",
		MongoDatabase: "malou",
	})
	t.Run("TestAddRunner", func(t *testing.T) {
		got, err := AddRunner(context.Background(), &Runner{
			Name: utils.StringWithCharsetV4(15),
			Labels: []string{
				"l" + utils.StringWithCharsetV1(1),
				"l" + utils.StringWithCharsetV1(1),
			},
		})
		if err != nil {
			t.Errorf("GetRunnerList() error = %v", err)
			return
		}
		t.Log(got)
	})
}

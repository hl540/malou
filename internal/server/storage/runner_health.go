package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DiskInfo struct {
	Total       uint64  `bson:"total"`
	Used        uint64  `bson:"used"`
	Free        uint64  `bson:"free"`
	UsedPercent float64 `bson:"used_percent"`
}

type RunnerHealth struct {
	RunnerID     string            `bson:"runner_id"`
	CreatedAt    int64             `bson:"created_at"`
	CpuPercent   []float64         `bson:"cpu_percent"`
	MemoryInfo   *DiskInfo         `bson:"memory_info"`
	DiskInfo     *DiskInfo         `bson:"disk_info"`
	WorkerStatus map[string]string `bson:"worker_status"`
}

func AddRunnerHealth(ctx context.Context, data *RunnerHealth) error {
	filter := bson.M{
		"runner_id": data.RunnerID,
	}
	opts := options.FindOne().SetSkip(9).SetSort(bson.M{"created_at": -1})
	result := RunnerHealthColl.FindOne(ctx, filter, opts)
	var health *RunnerHealth
	if err := result.Decode(&health); err == nil {
		delFilter := bson.M{
			"runner_id":  data.RunnerID,
			"created_at": bson.M{"$lte": health.CreatedAt},
		}
		RunnerHealthColl.DeleteMany(ctx, delFilter)
	}
	_, err := RunnerHealthColl.InsertOne(ctx, data)
	return err
}

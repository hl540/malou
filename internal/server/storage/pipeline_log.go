package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PipelineLog struct {
	PipelineID string `bson:"pipeline_id"`
	Step       string `bson:"step"`
	Cmd        string `bson:"cmd"`
	Message    string `bson:"message"`
	Type       string `bson:"type"`
	Timestamp  int64  `bson:"timestamp"`
	Duration   int64  `bson:"duration"`
}

func AddPipelineLog(ctx context.Context, data *PipelineLog) error {
	_, err := PipelineLogColl.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func GetPipelineLogs(ctx context.Context, pipelineID string, offset int64) ([]*PipelineLog, error) {
	filter := bson.M{
		"pipeline_id": pipelineID,
	}
	opts := options.Find().SetSkip(offset)
	cur, err := PipelineLogColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var results []*PipelineLog
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

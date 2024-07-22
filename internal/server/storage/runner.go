package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/hl540/malou/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Runner struct {
	ID     string   `bson:"id"`
	Secret string   `bson:"secret"`
	Name   string   `bson:"name"`
	Labels []string `bson:"labels"`
}

func AddRunner(ctx context.Context, name string, labels []string) (string, error) {
	runner := &Runner{
		ID:     uuid.New().String(),
		Secret: utils.StringWithCharsetV4(20),
		Name:   name,
		Labels: labels,
	}
	if _, err := RunnerColl.InsertOne(ctx, runner); err != nil {
		return "", err
	}
	return runner.ID, nil
}

type RunnerListOption struct {
	Name   string
	Labels []string
	Limit  int64
	Page   int64
}

func GetRunnerList(ctx context.Context, option *RunnerListOption) ([]*Runner, int64, error) {
	filter := bson.M{}
	if option.Name != "" {
		filter["name"] = bson.M{"$regex": option.Name}
	}
	if len(option.Labels) != 0 {
		filter["labels"] = bson.M{"$in": option.Labels}
	}

	// 查询总数
	count, err := RunnerColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return make([]*Runner, 0), 0, nil
	}

	// 分页
	skipValue := (option.Limit - 1) * option.Page
	opts := options.Find().SetSkip(skipValue).SetLimit(option.Limit)
	cur, err := RunnerColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}

	// 查询列表
	var results []*Runner
	if err := cur.All(ctx, &results); err != nil {
		return nil, 0, err
	}
	return results, count, nil
}

func GetRunnerByID(ctx context.Context, runnerID string) (*Runner, error) {
	var result Runner
	err := RunnerColl.FindOne(ctx, bson.D{{"id", runnerID}}).Decode(&result)
	return &result, err
}

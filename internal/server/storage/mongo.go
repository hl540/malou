package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var mongoClient *mongo.Client
var database *mongo.Database

const (
	RunnerCollection       = "runners"
	RunnerHealthCollection = "runners_health"
	PipelineLogCollection  = "pipeline_log"
)

var (
	RunnerColl       *mongo.Collection
	RunnerHealthColl *mongo.Collection
	PipelineLogColl  *mongo.Collection
)

func Transaction(ctx context.Context, fn func(ctx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)

	session, err := mongoClient.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	return session.WithTransaction(ctx, fn, txnOptions)
}

package storage

import (
	"context"
	"github.com/hl540/malou/internal/server"
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

func InitMongo(config *server.Config) (*mongo.Client, error) {
	var err error
	mongoClient, err = mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(config.MongoUri),
	)
	if err != nil {
		return nil, err
	}

	database = mongoClient.Database(config.MongoDatabase)
	RunnerColl = database.Collection(RunnerCollection)
	RunnerHealthColl = database.Collection(RunnerHealthCollection)
	PipelineLogColl = database.Collection(PipelineLogCollection)
	return mongoClient, nil
}

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

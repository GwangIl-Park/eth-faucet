package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"eth-faucet/internal/logger"

	log "github.com/sirupsen/logrus"
)

func ConnectDB(ctx context.Context, mongoUri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)

	logger.Logger.WithFields(log.Fields{
		"uri": mongoUri,
	}).Info("Connecting Mongo DB")

	// MongoDB 연결
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.Logger.WithError(err).Error("Mongo DB Connect Error")
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Logger.WithError(err).Error("Mongo DB Ping Error")
		return nil, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, dbName string, colName string) *mongo.Collection {
	logger.Logger.WithFields(log.Fields{
		"dbName":     dbName,
		"Collection": colName,
	}).Debug("Get Collection")

	return client.Database(dbName).Collection(colName)
}

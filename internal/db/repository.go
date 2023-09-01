package db

import (
	"context"
	"eth-faucet/internal/logger"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertDocument(client *mongo.Client, dbName string, colName string, document Collection) error {
	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": document.Address,
		"amount":  document.Amount,
	}).Debug("Insert Document")

	result, err := GetCollection(client, dbName, colName).InsertOne(context.TODO(), document)

	if err == nil {
		logger.Logger.WithFields(log.Fields{
			"colName":  colName,
			"insertId": result.InsertedID,
		}).Debug("Insert Document Success")
	}

	return err
}

func ReadDocument(client *mongo.Client, dbName string, colName string, address string) (*Collection, error) {
	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": address,
	}).Debug("Read Document")

	var result Collection

	filter := bson.M{"address": address}

	err := GetCollection(client, dbName, colName).FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": result.Address,
		"amount":  result.Amount,
		"Time":    result.Time,
	}).Debug("Read Document Result")

	return &result, nil
}

// Update func
func UpdateDocument(client *mongo.Client, dbName string, colName string, document Collection) error {
	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": document.Address,
		"amount":  document.Amount,
	}).Debug("Update Document")

	filter := bson.M{"address": document.Address}

	update := bson.M{
		"$set": document,
	}

	_, err := GetCollection(client, dbName, colName).UpdateOne(context.TODO(), filter, update)

	if err == nil {
		logger.Logger.WithFields(log.Fields{
			"colName": colName,
		}).Debug("Update Document Success")
	}

	return err
}

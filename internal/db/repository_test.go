package db_test

import (
	"context"
	"eth-faucet/internal/db"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryTestSuite struct {
	suite.Suite
	client *mongo.Client
	time   time.Time
}

func (ts *RepositoryTestSuite) SetupSuite() {
	ts.client, _ = db.ConnectDB(context.Background(), "mongodb://localhost:27017")
	ts.time = time.Now()
	filter := bson.M{"address": "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"}
	db.GetCollection(ts.client, "db_test", "col_test").DeleteOne(context.Background(), filter)
}

func (ts *RepositoryTestSuite) SetupTest() {
}

func (ts *RepositoryTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s\n", suiteName, testName)
}

func (ts *RepositoryTestSuite) TestInsertDocument() {
	err := db.InsertDocument(ts.client, "db_test", "col_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A", 1, ts.time)
	assert.NoError(ts.T(), err, "InsertDocument should not return error")
}

func (ts *RepositoryTestSuite) TestReadDocument() {
	document, err := db.ReadDocument(ts.client, "db_test", "col_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")

	assert.NoError(ts.T(), err, "ReadDocument should not return error")

	assert.Equal(ts.T(), document.Address, "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A", "ReadDocument address should be 0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	assert.Equal(ts.T(), document.Amount, uint64(1), "ReadDocument amount should be 1")
	//assert.Equal(ts.T(), document.Time.Local(), ts.time, "ReadDocument time should be %s", ts.time)
}

func (ts *RepositoryTestSuite) TestUpdateDocument() {
	updateTime := time.Now()
	err := db.UpdateDocument(ts.client, "db_test", "col_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A", 10, updateTime)

	assert.NoError(ts.T(), err, "UpdateDocument should not return error")

	document, err := db.ReadDocument(ts.client, "db_test", "col_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")

	assert.Equal(ts.T(), document.Address, "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A", "ReadDocument address should be 0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	assert.Equal(ts.T(), document.Amount, uint64(10), "ReadDocument amount should be 10")
	//assert.Equal(ts.T(), document.Time.Local(), ts.time, "ReadDocument time should be %s", updateTime)
}

func (ts *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s End\n", suiteName, testName)
}

// suite에서 각 테스트 실행후에 실행된다
func (ts *RepositoryTestSuite) TearDownTest() {
}

// suite에서 모든 테스트가 실행된 후에 실행된다
func (ts *RepositoryTestSuite) TearDownSuite() {
	filter := bson.M{"address": "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"}
	db.GetCollection(ts.client, "db_test", "col_test").DeleteOne(context.Background(), filter)
	fmt.Printf("mongo.go Test End\n")
	ts.client.Disconnect(context.Background())
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

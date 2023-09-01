package db_test

import (
	"context"
	"eth-faucet/internal/db"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MongoTestSuite struct {
	suite.Suite
}

func (ts *MongoTestSuite) SetupSuite() {
}

func (ts *MongoTestSuite) SetupTest() {
}

func (ts *MongoTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s\n", suiteName, testName)
}

func (ts *MongoTestSuite) TestConnectDB() {
	client, err := db.ConnectDB(context.Background(), "mongodb://localhost:27017")
	assert.NoError(ts.T(), err, "ConnectDB should return no err")
	client.Disconnect(context.Background())
}

func (ts *MongoTestSuite) TestGetCollection() {
	client, _ := db.ConnectDB(context.Background(), "mongodb://localhost:27017")
	collection := db.GetCollection(client, "faucet_test", "Token_test")
	assert.NotNil(ts.T(), collection, "GetCollection with Token should not nil")
	client.Disconnect(context.Background())
}

func (ts *MongoTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s End\n", suiteName, testName)
}

// suite에서 각 테스트 실행후에 실행된다
func (ets *MongoTestSuite) TearDownTest() {
}

// suite에서 모든 테스트가 실행된 후에 실행된다
func (ets *MongoTestSuite) TearDownSuite() {
	fmt.Printf("mongo.go Test End\n")
}

func TestMongo(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}

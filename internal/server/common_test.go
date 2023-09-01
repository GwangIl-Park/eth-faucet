package server_test

import (
	"context"
	"eth-faucet/internal/db"
	"eth-faucet/internal/geth"
	"eth-faucet/internal/server"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type CommonTestSuite struct {
	suite.Suite
	server *server.Server
}

func (ts *CommonTestSuite) SetupSuite() {
	ts.server = new(server.Server)
	ts.server.Db, _ = db.ConnectDB(context.Background(), "mongodb://localhost:27017")
	geth.InitGenesis("/../../config/genesis.json")
}

func (ts *CommonTestSuite) SetupTest() {
}

func (ts *CommonTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s\n", suiteName, testName)
}

func (ts *CommonTestSuite) TestCheckLimit() {
	filter := bson.M{"address": "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"}
	db.GetCollection(ts.server.Db, "db_common_test", "col_common_test").DeleteOne(context.Background(), filter)
	db.InsertDocument(ts.server.Db, "db_common_test", "col_common_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A", 1, time.Now())

	server.LIMIT_SECOND = 24 * 360

	ballow, _, err := ts.server.CheckLimit("db_common_test", "col_common_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	assert.NoError(ts.T(), err, "CheckLimit should not return error")
	assert.Equal(ts.T(), ballow, false, "ballow should be false in limit")

	time.Sleep(time.Second)

	server.LIMIT_SECOND = 1
	ballow, _, err = ts.server.CheckLimit("db_common_test", "col_common_test", "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	assert.Equal(ts.T(), ballow, true, "ballow should be true out limit")

	db.GetCollection(ts.server.Db, "db_common_test", "col_common_test").DeleteOne(context.Background(), filter)
}

func (ts *CommonTestSuite) TestSignAndSendTransaction() {
	to := geth.HexToAddress("0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	ts.server.Client, _ = ethclient.Dial("http://localhost:8545")
	chainId := big.NewInt(4693)
	privateKey, _ := geth.HexToECDSA("296ec1b0a6d29469c28b930c24358f24538d2562847f0505dd8d73d9b427f17a")

	sender := &server.SenderAccount{
		PrivateKey: privateKey,
		Address:    geth.PublicKeyToAddress(&privateKey.PublicKey),
	}

	ts.server.Common = &server.Common{
		sender,
		chainId,
	}
	tx, _ := ts.server.MakeTransaction(to, big.NewInt(0), nil)
	_, _, err := ts.server.SignAndSendTransaction(tx)
	assert.NoError(ts.T(), err, "SignAndSendTransaction should not return error")
}

func (ts *CommonTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s End\n", suiteName, testName)
}

// suite에서 각 테스트 실행후에 실행된다
func (ts *CommonTestSuite) TearDownTest() {
}

// suite에서 모든 테스트가 실행된 후에 실행된다
func (ts *CommonTestSuite) TearDownSuite() {
	fmt.Printf("common.go Test End\n")
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}

package server_test

import (
	"context"
	"eth-faucet/internal/config"
	"eth-faucet/internal/db"
	"eth-faucet/internal/server"
	faucetETH "eth-faucet/proto/faucetETH"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type ETHRequestTestSuite struct {
	suite.Suite
	server *server.Server
}

func (ts *ETHRequestTestSuite) SetupSuite() {
	ts.server = new(server.Server)
	cfg := config.Config{
		PrivateKeyHex:    "296ec1b0a6d29469c28b930c24358f24538d2562847f0505dd8d73d9b427f17a",
		EthereumURL:      "http://localhost:8545",
		EthAmount:        1,
		TokenAmount:      1,
		TokenAddress:     "0x04b268e462a54d6fce12a919c1d492ee13599b1b",
		GrpcServerUrl:    "localhost:5052",
		GatewayServerUrl: "localhost:5053",
		MongoUri:         "mongodb://localhost:27017",
		Limit:            10,
		LimitUnit:        "s",
	}
	ts.server.NewServer(&cfg)
	go ts.server.Start(&cfg)
}

func (ts *ETHRequestTestSuite) SetupTest() {
}

func (ts *ETHRequestTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s\n", suiteName, testName)
}

func (ts *ETHRequestTestSuite) TestRequestETH() {
	conn, _ := grpc.Dial("localhost:5052", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	client := faucetETH.NewFaucetETHClient(conn)

	address := common.HexToAddress("0x7d391b793e208f515a40c08dc87a2af1e53ffd9A")
	ethBefore, tokenBefore, _ := ts.server.GetEthTokenBalance(&address)

	_, err := client.RequestETH(context.Background(), &faucetETH.FaucetETHRequest{WalletAddress: "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"})

	assert.NoError(ts.T(), err, "RequestETH should not return error")

	ethAfter, tokenAfter, _ := ts.server.GetEthTokenBalance(&address)

	assert.Equal(ts.T(), ethBefore+1, ethAfter, "RequestETH ethBalance should equal")
	assert.Equal(ts.T(), tokenBefore, tokenAfter, "RequestETH TokenBalance should equal +1")

	_, err = client.RequestETH(context.Background(), &faucetETH.FaucetETHRequest{WalletAddress: "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"})
	assert.Error(ts.T(), err, "RequestETH in limit should return error")
}

func (ts *ETHRequestTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s End\n", suiteName, testName)
}

// suite에서 각 테스트 실행후에 실행된다
func (ts *ETHRequestTestSuite) TearDownTest() {
}

// suite에서 모든 테스트가 실행된 후에 실행된다
func (ts *ETHRequestTestSuite) TearDownSuite() {
	filter := bson.M{"address": "0x7d391b793e208f515a40c08dc87a2af1e53ffd9A"}
	db.GetCollection(ts.server.Db, "db_test", "col_test").DeleteOne(context.Background(), filter)
	ts.server.GrpcServer.GracefulStop()
	ts.server.GrpcGwServer.Close()
	fmt.Printf("processETHRequest.go Test End\n")
}

func TestETHRequest(t *testing.T) {
	suite.Run(t, new(ETHRequestTestSuite))
}

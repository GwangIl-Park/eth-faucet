package geth_test

import (
	"encoding/json"
	"eth-faucet/internal/geth"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GethTestSuite struct {
	suite.Suite
	client *ethclient.Client
}

func (ts *GethTestSuite) SetupSuite() {
	ts.client, _ = ethclient.Dial("http://localhost:8545")
}

func (ts *GethTestSuite) SetupTest() {
}

func (ts *GethTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s\n", suiteName, testName)
}

func (ts *GethTestSuite) TestInitGenesis() {
	assert.NoError(ts.T(), geth.InitGenesis("/../../config/genesis.json"), "With right path should return no err")
	assert.Error(ts.T(), geth.InitGenesis("/../../config/genesis_wrong.json"), "With wrong path should return err")
}

func (ts *GethTestSuite) TestGetBlockNumber() {
	blockNumber, err := geth.GetBlockNumber(ts.client)
	assert.NoError(ts.T(), err, "GetBlockNumber should return no err")
	assert.Greater(ts.T(), blockNumber, uint64(0), "GetBlockNumber should return greater than 0")
}

func (ts *GethTestSuite) TestGetChainID() {
	var genesis core.Genesis

	rootPath, _ := os.Getwd()
	genesisFile, err := os.Open(rootPath + "/../../config/genesis.json")

	defer genesisFile.Close()

	decoder := json.NewDecoder(genesisFile)
	err = decoder.Decode(&genesis)

	chainID, err := geth.GetChainID(ts.client)
	assert.NoError(ts.T(), err, "GetChainID should return no err")
	assert.Equal(ts.T(), chainID, genesis.Config.ChainID, "GetBlockNumber should return greater than 0")
}

func (ts *GethTestSuite) TestLeftPadBytes() {
	testBytes := []byte{248, 212, 93, 116, 76, 38, 206, 228, 188, 97, 109, 4, 22, 180, 221, 36, 95, 26, 7, 20}
	resultBytes := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 248, 212, 93, 116, 76, 38, 206, 228, 188, 97, 109, 4, 22, 180, 221, 36, 95, 26, 7, 20}

	assert.Equal(ts.T(), geth.LeftPadBytes(testBytes, 32), resultBytes, "LeftPadBytes is difference")
}

func (ts *GethTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("Test :: suiteName: %s, testName: %s End\n", suiteName, testName)
}

// suite에서 각 테스트 실행후에 실행된다
func (ets *GethTestSuite) TearDownTest() {
}

// suite에서 모든 테스트가 실행된 후에 실행된다
func (ts *GethTestSuite) TearDownSuite() {
	fmt.Printf("geth.go Test End\n")
}

func TestGeth(t *testing.T) {
	suite.Run(t, new(GethTestSuite))
}

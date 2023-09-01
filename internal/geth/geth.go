package geth

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"eth-faucet/internal/logger"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var genesis core.Genesis

func InitGenesis(path string) error {
	rootPath, _ := os.Getwd()
	genesisFile, err := os.Open(rootPath + path)

	if err != nil {
		logger.Logger.WithError(err).Error("GenesisFile Open Error")
		return err
	}

	defer genesisFile.Close()

	decoder := json.NewDecoder(genesisFile)
	err = decoder.Decode(&genesis)

	if err != nil {
		logger.Logger.WithError(err).Error("GenesisFile Decode Error")
		return err
	}

	return nil
}

//ethclient

func GetBlockNumber(client *ethclient.Client) (uint64, error) {
	return client.BlockNumber(context.Background())
}

func GetChainID(client *ethclient.Client) (*big.Int, error) {
	return client.ChainID(context.Background())
}

func GetNonce(client *ethclient.Client, account *common.Address) (uint64, error) {
	return client.PendingNonceAt(context.Background(), *account)
}

func GetBalance(client *ethclient.Client, account *common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), *account, nil)
}

func CallContract(client *ethclient.Client, data []byte, to *common.Address) ([]byte, error) {
	return client.CallContract(context.Background(), ethereum.CallMsg{Data: data, To: to}, nil)
}

func SendTransaction(client *ethclient.Client, tx *types.Transaction) error {
	return client.SendTransaction(context.Background(), tx)
}

//common

func LeftPadBytes(slice []byte, l int) []byte {
	return common.LeftPadBytes(slice, 32)
}

func HexToAddress(addressHex string) *common.Address {
	address := common.HexToAddress(addressHex)
	return &address
}

//crypto

func Keccak256(message []byte) []byte {
	return crypto.Keccak256(message)
}

func HexToECDSA(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKeyHex)
}

func PublicKeyToAddress(publicKey *ecdsa.PublicKey) *common.Address {
	address := crypto.PubkeyToAddress(*publicKey)
	return &address
}

//types

func GetSigner(chainId *big.Int, client *ethclient.Client) (types.Signer, error) {
	blockNumber, err := GetBlockNumber(client)
	if err != nil {
		return nil, err
	}

	return types.MakeSigner(genesis.Config, big.NewInt(int64(blockNumber))), nil
}

func SignTx(privateKey *ecdsa.PrivateKey, s types.Signer, txdata types.TxData) (*types.Transaction, error) {
	return types.SignNewTx(privateKey, s, txdata)
}

//bind

func WaitMined(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(context.Background(), client, tx)
}

package server

import (
	"crypto/ecdsa"
	"eth-faucet/internal/db"
	"eth-faucet/internal/geth"
	"eth-faucet/internal/logger"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	TRANSFER_SELECTOR  = geth.Keccak256([]byte("transfer(address,uint256)"))[:4]
	BALANCEOF_SELECTOR = geth.Keccak256([]byte("balanceOf(address)"))[:4]
	LIMIT_SECOND       float64
)

type SenderAccount struct {
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
}

type Common struct {
	Sender  *SenderAccount
	ChainId *big.Int
}

func (server *Server) SetCommon() error {
	privateKey, err := geth.HexToECDSA(server.Config.PrivateKeyHex)
	if err != nil {
		return err
	}

	sender := &SenderAccount{
		PrivateKey: privateKey,
		Address:    geth.PublicKeyToAddress(&privateKey.PublicKey),
	}

	chainId, err := geth.GetChainID(server.Client)
	if err != nil {
		return err
	}

	server.Common = &Common{
		sender,
		chainId,
	}

	return nil
}

func SetLimitSecond(limit float64, limitUnit byte) {
	LIMIT_SECOND = float64(limit)
	switch limitUnit {
	case 'm':
		LIMIT_SECOND *= 60
		break
	case 'h':
		LIMIT_SECOND *= 360
		break
	}
}

func (server *Server) CheckLimit(dbName string, colName string, address string) (bool, *db.Collection, error) {
	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": address,
	}).Debug("Check Limit")

	document, err := db.ReadDocument(server.Db, dbName, colName, address)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"address": address,
			}).Debug("First Request")

			return true, nil, nil
		}
		logger.Logger.WithError(err).Error("Read Error")

		return true, nil, err
	}

	diff := time.Now().Sub(document.Time)

	if diff.Seconds() < LIMIT_SECOND {
		return false, document, nil
	}
	return true, document, nil
}

func (server *Server) CheckLimitAll(dbName string, colName string, address string) (bool, *db.CollectionAll, error) {
	logger.Logger.WithFields(log.Fields{
		"colName": colName,
		"address": address,
	}).Debug("Check Limit")

	document, err := db.ReadAllDocument(server.Db, dbName, colName, address)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"address": address,
			}).Debug("First Request")

			return true, nil, nil
		}
		logger.Logger.WithError(err).Error("Read Error")

		return true, nil, err
	}

	diff := time.Now().Sub(document.Time)

	if diff.Seconds() < LIMIT_SECOND {
		return false, document, nil
	}
	return true, document, nil
}

func (server *Server) MakeTransaction(to *common.Address, value *big.Int, data []byte) (*types.DynamicFeeTx, error) {
	nonce, err := geth.GetNonce(server.Client, server.Common.Sender.Address)
	if err != nil {
		logger.Logger.WithError(err).Error("GetNonce Error")

		return nil, err
	}

	gasFeeCap := big.NewInt(10)

	gas := uint64(1000000)

	return &types.DynamicFeeTx{
		ChainID:   server.Common.ChainId,
		Nonce:     nonce,
		To:        to,
		Value:     value,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		Data:      data}, nil
}

func (server *Server) SignAndSendTransaction(tx types.TxData) (*types.Transaction, *types.Receipt, error) {
	signer, err := geth.GetSigner(server.Common.ChainId, server.Client)
	if err != nil {
		return nil, nil, err
	}

	signedTx, err := geth.SignTx(server.Common.Sender.PrivateKey, signer, tx)

	if err != nil {
		logger.Logger.WithError(err).Error("Sign transaction Error")

		return nil, nil, err
	}

	err = geth.SendTransaction(server.Client, signedTx)

	if err != nil {
		logger.Logger.WithError(err).Error("Send transaction Error")

		return nil, nil, err
	}

	receipt, err := geth.WaitMined(server.Client, signedTx)

	if err != nil {
		logger.Logger.WithError(err).Error("Wait Confirm Error")

		return nil, nil, err
	}

	return signedTx, receipt, nil
}

func (server *Server) UpdateDB(colName string, toAddressString string, amount string) {
	document, _ := db.ReadDocument(server.Db, "faucet", colName, toAddressString)
	if document == nil {
		err := db.InsertDocument(server.Db, "faucet", colName, db.Collection{Address: toAddressString, Amount: amount, Time: time.Now()})
		if err != nil {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"error":   err,
			}).Warn("Insert Document Error")
		}
	} else {
		currentAmount := new(big.Int)
		currentAmount.SetString(document.Amount, 10)
		addAmount := new(big.Int)
		addAmount.SetString(amount, 10)
		newAmount := new(big.Int)
		newAmount.Add(currentAmount, addAmount)

		err := db.UpdateDocument(server.Db, "faucet", colName, db.Collection{Address: toAddressString, Amount: newAmount.String(), Time: time.Now()})
		if err != nil {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"error":   err,
			}).Warn("Update Document Error")
		}
	}
}

func (server *Server) UpdateAllDB(colName string, toAddressString string, amountETH string, amountToken string) {
	document, _ := db.ReadAllDocument(server.Db, "faucet", colName, toAddressString)
	if document == nil {
		err := db.InsertAllDocument(server.Db, "faucet", colName, db.CollectionAll{Address: toAddressString, AmountETH: amountETH, AmountToken: amountToken, Time: time.Now()})
		if err != nil {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"error":   err,
			}).Warn("Insert Document Error")
		}
	} else {
		currentAmountETH := new(big.Int)
		currentAmountETH.SetString(document.AmountETH, 10)
		addAmountETH := new(big.Int)
		addAmountETH.SetString(amountETH, 10)
		newAmountETH := new(big.Int)
		newAmountETH.Add(currentAmountETH, addAmountETH)

		currentAmountToken := new(big.Int)
		currentAmountToken.SetString(document.AmountToken, 10)
		addAmountToken := new(big.Int)
		addAmountToken.SetString(amountToken, 10)
		newAmountToken := new(big.Int)
		newAmountToken.Add(currentAmountToken, addAmountToken)

		err := db.UpdateAllDocument(server.Db, "faucet", colName, db.CollectionAll{Address: toAddressString, AmountETH: newAmountETH.String(), AmountToken: newAmountToken.String(), Time: time.Now()})
		if err != nil {
			logger.Logger.WithFields(log.Fields{
				"colName": colName,
				"error":   err,
			}).Warn("Update Document Error")
		}
	}
}

func (server *Server) GetEthTokenBalance(to *common.Address) (*big.Int, *big.Int, error) {
	ethBalance, err := geth.GetBalance(server.Client, to)
	if err != nil {
		return nil, nil, err
	}

	paddedAddress := geth.LeftPadBytes(to.Bytes(), 32)

	var balanceData []byte
	balanceData = append(balanceData, BALANCEOF_SELECTOR...)
	balanceData = append(balanceData, paddedAddress...)

	result, err := geth.CallContract(server.Client, balanceData, geth.HexToAddress(server.Config.TokenAddress))
	if err != nil {
		return nil, nil, err
	}

	tokenBalance := big.NewInt(0).SetBytes(result)

	return ethBalance, tokenBalance, nil
}

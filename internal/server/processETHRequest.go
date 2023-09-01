package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"eth-faucet/internal/geth"
	faucetETH "eth-faucet/proto/faucetETH"

	"eth-faucet/internal/logger"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

func (server *Server) ProcessETH(address *common.Address) (string, error) {
	ethAmount := new(big.Int)
	ethAmount.SetString(server.Config.EthAmount, 10)

	tx, err := server.MakeTransaction(address, ethAmount, nil)
	if err != nil {
		logger.Logger.WithError(err).Error("MakeTransaction Fail")
		return "", errors.New(fmt.Errorf("MakeTransaction Fail : %s", err.Error()).Error())
	}

	signedTx, receipt, err := server.SignAndSendTransaction(tx)
	if err != nil {
		logger.Logger.WithError(err).Error("SignAndSendTransaction Fail")
		return "", errors.New(fmt.Errorf("SignAndSendTransaction Fail : %s", err.Error()).Error())
	}

	if receipt.Status == 0 {
		logger.Logger.Error("Send Transaction Status 0")
		return "", errors.New("Send Transaction Status 0")
	}

	txHash := signedTx.Hash().String()

	logger.Logger.WithFields(log.Fields{
		"account": address.String(),
		"txHash":  txHash,
	}).Info("Send ETH Success")

	return txHash, nil
}

func (server *Server) RequestETH(ctx context.Context, req *faucetETH.FaucetETHRequest) (*faucetETH.FaucetETHResponse, error) {
	toAddress := geth.HexToAddress(req.GetWalletAddress())

	bAllow, document, err := server.CheckLimit("faucet", "ETH", toAddress.String())
	if err != nil {
		return nil, err
	}

	if !bAllow {
		logger.Logger.WithFields(log.Fields{
			"address": document.Address,
			"Time":    document.Time,
		}).Debug("ETH Not Allow Yet")
		return nil, errors.New("You Can't Request ETH Yet")
	}

	txHash, err := server.ProcessETH(toAddress)

	if err != nil {
		logger.Logger.WithError(err).Error("RequestETH Fail")
		return nil, errors.New(fmt.Errorf("RequestToken Fail : %s", err.Error()).Error())
	}

	server.UpdateDB("ETH", toAddress.String(), server.Config.EthAmount)

	ethBalance, tokenBalance, err := server.GetEthTokenBalance(toAddress)
	if err != nil {
		logger.Logger.Warn("GetEthTokenBalance Fail")
		return nil, errors.New(fmt.Errorf("GetEthTokenBalance Fail : %s", err.Error()).Error())
	}

	return &faucetETH.FaucetETHResponse{TransactionHash: txHash, EthBalance: ethBalance.String(), TokenBalance: tokenBalance.String()}, nil
}

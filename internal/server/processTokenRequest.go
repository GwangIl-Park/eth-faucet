package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	faucetToken "eth-faucet/proto/faucetToken"

	"eth-faucet/internal/geth"
	"eth-faucet/internal/logger"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

func (server *Server) ProcessToken(address *common.Address) (string, error) {
	tokenAmount := new(big.Int)
	tokenAmount.SetString(server.Config.TokenAmount, 10)

	paddedAddress := geth.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := geth.LeftPadBytes(tokenAmount.Bytes(), 32)
	var data []byte
	data = append(data, TRANSFER_SELECTOR...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx, err := server.MakeTransaction(geth.HexToAddress(server.Config.TokenAddress), big.NewInt(0), data)
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
	}).Info("Send Token Success")

	return txHash, nil
}

func (server *Server) RequestToken(ctx context.Context, req *faucetToken.FaucetTokenRequest) (*faucetToken.FaucetTokenResponse, error) {
	toAddress := geth.HexToAddress(req.GetWalletAddress())

	bAllow, document, err := server.CheckLimit("faucet", "Token", toAddress.String())
	if err != nil {
		return nil, err
	}

	if !bAllow {
		logger.Logger.WithFields(log.Fields{
			"address": document.Address,
			"Time":    document.Time,
		}).Debug("Token Not Allow Yet")
		return nil, errors.New("You Can't Request Token Yet")
	}

	txHash, err := server.ProcessToken(toAddress)

	if err != nil {
		logger.Logger.WithError(err).Error("RequestToken Fail")
		return nil, errors.New(fmt.Errorf("RequestToken Fail : %s", err.Error()).Error())
	}

	server.UpdateDB("Token", toAddress.String(), server.Config.TokenAmount)

	ethBalance, TokenBalance, err := server.GetEthTokenBalance(toAddress)
	if err != nil {
		logger.Logger.Warn("GetEthTokenBalance Fail")
		return nil, errors.New(fmt.Errorf("GetEthTokenBalance Fail : %s", err.Error()).Error())
	}

	return &faucetToken.FaucetTokenResponse{TransactionHash: txHash, EthBalance: ethBalance.String(), TokenBalance: TokenBalance.String()}, nil
}

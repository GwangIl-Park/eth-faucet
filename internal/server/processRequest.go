package server

import (
	"context"
	"errors"
	"fmt"

	"eth-faucet/internal/geth"
	"eth-faucet/internal/logger"
	faucet "eth-faucet/proto/faucet"

	log "github.com/sirupsen/logrus"
)

func (server *Server) Request(ctx context.Context, req *faucet.FaucetRequest) (*faucet.FaucetResponse, error) {
	toAddressString := req.GetWalletAddress()
	toAddress := geth.HexToAddress(toAddressString)

	chResult := make(chan QueueResponse)

	server.Enqueue(toAddress, chResult)

	response := <-chResult

	return response.response, response.err
}

func (server *Server) ProcessQueue(element Element) {
	toAddress := element.Address

	bAllow, document, err := server.CheckLimitAll("faucet", "All", toAddress.String())
	if err != nil {
		element.chResult <- QueueResponse{nil, err}
		return
	}

	if !bAllow {
		logger.Logger.WithFields(log.Fields{
			"address": document.Address,
			"Time":    document.Time,
		}).Debug("Not Allow Yet")
		element.chResult <- QueueResponse{nil, errors.New("You Can't Request Yet")}
		return
	}

	logger.Logger.WithFields(log.Fields{
		"address": toAddress.String(),
	}).Debug("Request")

	ethTxHash, err := server.ProcessETH(toAddress)
	if err != nil {
		element.chResult <- QueueResponse{nil, errors.New(fmt.Errorf("Both ETH and Token Fail : %s", err.Error()).Error())}
		return
	}

	tokenTxHash, err := server.ProcessToken(toAddress)
	if err != nil {
		element.chResult <- QueueResponse{nil, errors.New(fmt.Errorf("ETH Success but Token Fail : %s", err.Error()).Error())}
		return
	}

	server.UpdateAllDB("All", toAddress.String(), server.Config.EthAmount, server.Config.TokenAmount)

	ethBalance, tokenBalance, err := server.GetEthTokenBalance(toAddress)

	element.chResult <- QueueResponse{&faucet.FaucetResponse{EthTransactionHash: ethTxHash, TokenTransactionHash: tokenTxHash, EthBalance: ethBalance.String(), TokenBalance: tokenBalance.String()}, nil}
}

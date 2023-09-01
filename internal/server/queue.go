package server

import (
	"eth-faucet/internal/logger"
	faucet "eth-faucet/proto/faucet"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type QueueResponse struct {
	response *faucet.FaucetResponse
	err      error
}

type Element struct {
	Address  *common.Address
	chResult chan QueueResponse
}

type Queue struct {
	Elements []Element
}

func (server *Server) Enqueue(address *common.Address, chResult chan QueueResponse) {
	logger.Logger.WithFields(log.Fields{
		"address": address,
		"length":  len(server.Queue.Elements),
	}).Debug("Enqueue")
	server.Queue.Elements = append(server.Queue.Elements, Element{address, chResult})
}

func (server *Server) Dequeue() {
	for {
		length := len(server.Queue.Elements)
		if length != 0 {
			element := server.Queue.Elements[0]
			logger.Logger.WithFields(log.Fields{
				"address": element.Address,
				"length":  length,
			}).Debug("Dequeue")
			server.Queue.Elements = server.Queue.Elements[1:]
			server.ProcessQueue(element)
		}
	}
}

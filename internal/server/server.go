package server

import (
	"context"
	"eth-faucet/internal/db"
	faucet "eth-faucet/proto/faucet"
	faucetETH "eth-faucet/proto/faucetETH"
	faucetToken "eth-faucet/proto/faucetToken"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"

	"eth-faucet/internal/config"

	"google.golang.org/grpc"
)

type Server struct {
	faucet.UnimplementedFaucetServer
	faucetETH.UnimplementedFaucetETHServer
	faucetToken.UnimplementedFaucetTokenServer
	Db           *mongo.Client
	Config       *config.Config
	Client       *ethclient.Client
	Common       *Common
	GrpcServer   *grpc.Server
	GrpcGwServer *http.Server
	Queue        *Queue
}

func (server *Server) NewServer(cfg *config.Config) error {
	server.Config = cfg

	var err error

	server.Client, err = ethclient.Dial(cfg.EthereumURL)

	if err != nil {
		return err
	}

	server.Db, err = db.ConnectDB(context.Background(), cfg.MongoUri)
	if err != nil {
		return err
	}

	SetLimitSecond(float64(cfg.Limit), cfg.LimitUnit[0])

	server.SetCommon()

	var q *Queue = new(Queue)

	server.Queue = q

	return nil
}

func (server *Server) Start(cfg *config.Config) error {
	chErr := make(chan error)
	go StartGrpcServer(cfg, server, chErr)

	go StartGrpcGateway(cfg, server, chErr)

	go server.Dequeue()

	err := <-chErr
	if err != nil {
		return err
	}

	return nil
}

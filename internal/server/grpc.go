package server

import (
	"context"
	"eth-faucet/internal/config"
	"eth-faucet/internal/logger"
	faucet "eth-faucet/proto/faucet"
	faucetETH "eth-faucet/proto/faucetETH"
	faucetToken "eth-faucet/proto/faucetToken"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func StartGrpcServer(cfg *config.Config, server *Server, chErr chan error) {
	lis, err := net.Listen("tcp", cfg.GrpcServerUrl)
	if err != nil {
		logger.Logger.WithError(err).Error("GRPC Server Listen Error")
		chErr <- err
	}
	server.GrpcServer = grpc.NewServer()
	faucet.RegisterFaucetServer(server.GrpcServer, server)
	faucetETH.RegisterFaucetETHServer(server.GrpcServer, server)
	faucetToken.RegisterFaucetTokenServer(server.GrpcServer, server)

	logger.Logger.WithFields(log.Fields{
		"GrpcServerUrl": cfg.GrpcServerUrl,
	}).Info("GRPC Server Start")

	if err := server.GrpcServer.Serve(lis); err != nil {
		logger.Logger.WithError(err).Error("GRPC Server Start Error")
		chErr <- err
	}
	defer server.GrpcServer.Stop()
}

func StartGrpcGateway(cfg *config.Config, server *Server, chErr chan error) {
	conn, err := grpc.Dial(cfg.GrpcServerUrl, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		logger.Logger.WithError(err).Error("GRPC Server Dial Error")
		chErr <- err
	}

	gwmux := runtime.NewServeMux()

	faucet.RegisterFaucetHandler(context.Background(), gwmux, conn)
	faucetETH.RegisterFaucetETHHandler(context.Background(), gwmux, conn)
	faucetToken.RegisterFaucetTokenHandler(context.Background(), gwmux, conn)

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(gwmux)

	server.GrpcGwServer = &http.Server{
		Addr:    cfg.GatewayServerUrl,
		Handler: withCors,
	}

	logger.Logger.WithFields(log.Fields{
		"GatewayServerUrl": cfg.GatewayServerUrl,
	}).Info("Gateway Server Start")

	err = server.GrpcGwServer.ListenAndServe()

	if err != nil {
		logger.Logger.WithError(err).Error("GRPC Gateway Start Error")
		chErr <- err
	}
	defer server.GrpcGwServer.Close()
}

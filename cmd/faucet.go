/*
Copyright © 2023 NAME HERE rkasud0@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"eth-faucet/internal/config"
	"eth-faucet/internal/geth"
	"eth-faucet/internal/logger"
	"eth-faucet/internal/server"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use: "faucet",
		RunE: func(command *cobra.Command, args []string) error {
			err := logger.SetLogger(verbosity)

			var cfg *config.Config
			if err := viper.Unmarshal(&cfg); err != nil {
				log.Fatal(err)
			}

			if err != nil {
				fmt.Println("SetLogger Error : ", err)
				os.Exit(1)
			}

			logger.Logger.WithFields(log.Fields{
				"privateKey":       cfg.PrivateKeyHex,
				"ethereumUrl":      cfg.EthereumURL,
				"ethSendAmount":    cfg.EthAmount,
				"tokenSendAmount":  cfg.TokenAmount,
				"tokenAddress":     cfg.TokenAddress,
				"grpcServerUrl":    cfg.GrpcServerUrl,
				"gatewayServerUrl": cfg.GatewayServerUrl,
				"mongoUri":         cfg.MongoUri,
				"limit":            cfg.Limit,
				"limitUnit":        cfg.LimitUnit,
				"verbosity":        verbosity,
			}).Debug("Check Flag")

			if cfg.PrivateKeyHex == "" {
				logger.Logger.Error("privateKey is required")
				os.Exit(1)
			}
			if cfg.TokenAddress == "" {
				logger.Logger.Error("tokenAddress is required")
				os.Exit(1)
			}

			genesisPath := "/config/genesis.json"
			err = geth.InitGenesis(genesisPath)
			if err != nil {
				logger.Logger.WithError(err).Error("Init Genesis Error")
				return err
			}

			var s *server.Server = new(server.Server)
			err = s.NewServer(cfg)

			if err != nil {
				logger.Logger.WithError(err).Error("New Server Error")
				return err
			}

			if err := s.Start(cfg); err != nil {
				logger.Logger.WithError(err).Error("Server Start Error")
				return err
			}
			return nil
		},
	}
)

var (
	verbosity string
)

func init() {
	rootCmd.Flags().String("privateKey", "", "Faucet Sender privateKey")
	rootCmd.Flags().String("ethereumUrl", "http://localhost:8545", "Ethereum Node URL")
	rootCmd.Flags().String("ethSendAmount", "1", "ETH Send Amount")
	rootCmd.Flags().String("tokenSendAmount", "1", "Token Send Amount")
	rootCmd.Flags().String("tokenAddress", "", "Token Address")
	rootCmd.Flags().String("grpcServerUrl", "192.168.10.52:5050", "GRPC Server URL")
	rootCmd.Flags().String("gatewayServerUrl", "192.168.10.52:5051", "Gateway Server URL")
	rootCmd.Flags().String("mongoUri", "mongodb://localhost:27017", "Mongo DB URI")
	rootCmd.Flags().Uint16("limit", 24, "Faucet Limit Hour")
	rootCmd.Flags().String("limitUnit", "h", "Faucet Limit Unit, s:second, m:minute, h:hour")
	rootCmd.Flags().StringVar(&verbosity, "verbosity", "info", "Verbosity Level")

	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Could not execute root command")
		os.Exit(1)
	}
}

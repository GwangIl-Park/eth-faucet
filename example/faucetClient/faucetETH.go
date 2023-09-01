package faucetclient

import (
	"context"
	"eth-faucet/proto/faucetETH"
	"fmt"

	"google.golang.org/grpc"
)

func RequestETH(address string) *faucetETH.FaucetETHResponse {
	fmt.Println("start")
	conn, _ := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	client := faucetETH.NewFaucetETHClient(conn)

	fmt.Println("ctx")
	res, err := client.RequestETH(context.Background(), &faucetETH.FaucetETHRequest{WalletAddress: address})

	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println("res")
	return res
}

package faucetclient

import (
	"context"
	"eth-faucet/proto/faucetToken"
	"fmt"

	"google.golang.org/grpc"
)

func RequestToken(address string) *faucetToken.FaucetTokenResponse {
	fmt.Println("start")
	conn, _ := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	client := faucetToken.NewFaucetTokenClient(conn)

	fmt.Println("ctx")
	res, err := client.RequestToken(context.Background(), &faucetToken.FaucetTokenRequest{WalletAddress: address})

	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println("res")
	return res
}

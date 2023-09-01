package main

import (
	faucetClient "eth-faucet/example/faucetClient"
	"fmt"
)

func main() {

	res := faucetClient.RequestETH("0xe95391ac993547bb3006c79618fcd28ab97377fc")
	fmt.Println(res.TransactionHash)
}

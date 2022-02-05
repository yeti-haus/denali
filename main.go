package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ethUrl = "http://127.0.0.1:8545"
const verbose = false

func main() {
	start := time.Now()
	ctx := context.Background()

	fmt.Println("We're go-ing now")

	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		panic(err)
	}

	startBlock := int64(6322283)

	for i := range make([]struct{}, 10000) {
		i := startBlock - int64(i)
		cb := big.NewInt(i)

		_, err := client.BlockByNumber(ctx, cb)
		if err != nil {
			panic(err)
		}

		_, err = client.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: cb, ToBlock: cb})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Finished in " + fmt.Sprint(time.Since(start)))
}

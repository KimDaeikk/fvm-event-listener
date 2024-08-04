package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func listenEvent(client *ethclient.Client, contractAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println("Error:", err)
			return
		case vLog := <-logs:
			fmt.Println("Log received:", vLog)
		}
	}
}

func main() {
	client, err := ethclient.Dial("http://222.121.167.136:1234/rpc/v1")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	contractAddresses := []string{
		"0xYourContractAddress1",
		"0xYourContractAddress2",
		// 추가 계약 주소
	}

	for _, addr := range contractAddresses {
		wg.Add(1)
		go listenEvent(client, addr, &wg)
	}

	wg.Wait()
}

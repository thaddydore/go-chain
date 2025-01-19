package events

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenToEvents(client *ethclient.Client, address string) {
    contractAddress := common.HexToAddress(address)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatalf("Failed to subscribe to logs: %v", err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatalf("Subscription error: %v", err)
        case vLog := <-logs:
            log.Printf("New Log: %+v", vLog)
        }
    }
}

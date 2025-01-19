package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thaddydore/gochain/contracts"
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum client: %v", err)
    }

    abiJSON := `...` // ABI JSON string of the contract
    contractAddress := "0xYourContractAddress"
    contract, err := contracts.NewContract(client, contractAddress, abiJSON)
    if err != nil {
        log.Fatalf("Failed to create contract: %v", err)
    }

    var result string
    err = contract.CallMethod("methodName", &result)
    if err != nil {
        log.Fatalf("Failed to call method: %v", err)
    }

    log.Printf("Result: %s", result)
}

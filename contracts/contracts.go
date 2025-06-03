package contracts

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Contract represents an Ethereum contract
type Contract struct {
    Address common.Address
    ABI     abi.ABI
    Client  *ethclient.Client
}

// NewContract initializes a new contract instance
func NewContract(client *ethclient.Client, address string, abiJSON string) (*Contract, error) {
    contractAddress := common.HexToAddress(address)
    parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
    if err != nil {
        return nil, fmt.Errorf("failed to parse ABI JSON: %w", err)
    }
    return &Contract{
        Address: contractAddress,
        ABI:     parsedABI,
        Client:  client,
    }, nil
}

// CallMethod calls a read-only method on the contract
func (c *Contract) CallMethod(methodName string, result interface{}, params ...interface{}) error {
    callData, err := c.ABI.Pack(methodName, params...)
    if err != nil {
        return fmt.Errorf("failed to pack method %q with params %v: %w", methodName, params, err)
    }
    res, err := c.Client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &c.Address,
        Data: callData,
    }, nil)
    if err != nil {
        return fmt.Errorf("failed to call contract method %q: %w", methodName, err)
    }
	// Unpack the raw response data into the provided result interface.
     // Adds method-specific context to the error if unpacking fails.
    if err := c.ABI.UnpackIntoInterface(result, methodName, res); err != nil {
		return fmt.Errorf("failed to unpack result for method %q: %w", methodName, err)
	}
	
	//this function completed successfully,no error occurred.
	return nil
}

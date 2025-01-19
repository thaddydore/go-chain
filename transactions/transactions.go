package transactions

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignTransaction signs a transaction with the private key
func SignTransaction(privateKeyHex string, tx *bind.TransactOpts) error {
    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return err
    }
    tx.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
        signer := types.LatestSignerForChainID(tx.ChainId())
        signedTx, err := types.SignTx(tx, signer, privateKey)
        return signedTx, err
    }
    return nil
}

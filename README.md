# Go-Chain

Go-based library for interacting with blockchain networks. It simplifies smart contract interactions, transaction building, event listening, and account management, primarily for Ethereum and Ethereum-compatible networks.

### Features

- Smart Contract Interaction:
  - Read/write methods for Ethereum smart contracts.
- Transaction Builder:
  - Build, sign, and send transactions.
- Event Listener:
  - Subscribe and listen to blockchain events in real-time.
- Account Management:
   - Generate, manage, and use wallets and private keys.
- Extendable:
    - Can be adapted for other EVM-compatible networks like Binance Smart Chain, Polygon, etc.

 ### Installation
To install the library, use:

```bash
go get [github.com/your-username/playgrounds-go](https://github.com/thaddydore/go-chain)
```

## Usage

### 1. Initialize a Blockchain Client

Connect to an Ethereum client (e.g., Infura, Alchemy, or a local node):


```bash
import (
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum client: %v", err)
    }
    // Use the client for contract interaction, transactions, etc.
}
```

### 2. Smart Contract Interaction

Call a smart contract method:

```bash

import (
    "log"

    "github.com/thaddydore/go-chain/contracts"
)

func main() {
    client, _ := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
    abiJSON := `...` // Your contract's ABI JSON string
    contractAddress := "0xYourContractAddress"

    contract, _ := contracts.NewContract(client, contractAddress, abiJSON)

    var result string
    err := contract.CallMethod("methodName", &result)
    if err != nil {
        log.Fatalf("Failed to call method: %v", err)
    }

    log.Printf("Result: %s", result)
}

```

### 3. Build and Sign Transactions

Create and sign transactions:

```bash

import (
    "github.com/thaddydore/go-chain/transactions"
)

func main() {
    privateKey := "0xYourPrivateKey"

    tx, _ := transactions.BuildTransaction(
        "0xRecipientAddress",
        "1000000000000000000", // 1 ETH in wei
        "0xContractData",
    )

    err := transactions.SignTransaction(privateKey, tx)
    if err != nil {
        log.Fatalf("Failed to sign transaction: %v", err)
    }

    // Send the signed transaction...
}

```

### 4. Listen to Blockchain Events
Subscribe to smart contract events:

```bash
import (
    "github.com/thaddydore/go-chain/events"
)

func main() {
    client, _ := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
    contractAddress := "0xYourContractAddress"

    events.ListenToEvents(client, contractAddress)
}

```

## Project Structure

```bash

playgrounds-go/
├── contracts/       # Smart contract interaction logic
├── transactions/    # Transaction creation and signing
├── events/          # Blockchain event listeners
├── accounts/        # Wallet and private key management
└── main.go          # Example usage

```

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue for suggestions or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Disclaimer

This library is provided "as is" and without warranty. Cryptocurrency interactions are inherently risky. Use this library at your own discretion, and ensure compliance with all applicable regulations.










 


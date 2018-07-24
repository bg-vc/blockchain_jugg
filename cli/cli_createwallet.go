package cli

import (
	"fmt"
	"blockchain_jugg/wallet"
)

func (cli *CLI) CreateWallet(nodeID string) {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your wallet address: %s\n", address)
}



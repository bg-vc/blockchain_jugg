package cli

import (
	"fmt"
	"github.com/bg-vc/blockchain_jugg/wallet"
	"log"
)

func (cli *CLI) ListAddresses(nodeID string) {
	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

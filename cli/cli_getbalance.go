package cli

import (
	"fmt"
	"github.com/bg-vc/blockchain_jugg/core"
	"github.com/bg-vc/blockchain_jugg/crypto"
	"github.com/bg-vc/blockchain_jugg/wallet"
	"log"
)

func (cli *CLI) GetBalance(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := crypto.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

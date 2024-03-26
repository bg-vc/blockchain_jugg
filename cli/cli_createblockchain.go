package cli

import (
	"fmt"
	"github.com/bg-vc/blockchain_jugg/core"
	"github.com/bg-vc/blockchain_jugg/wallet"
	"log"
)

func (cli *CLI) CreateBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}

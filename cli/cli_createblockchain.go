package cli


import (
	"fmt"
	"log"
	"blockchain_jugg/wallet"
	"blockchain_jugg/core"
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

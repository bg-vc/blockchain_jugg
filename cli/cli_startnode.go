package cli


import (
	"fmt"
	"log"
	"blockchain_jugg/wallet"
	"blockchain_jugg/p2p"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	p2p.StartServer(nodeID, minerAddress)
}

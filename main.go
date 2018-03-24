package main

func main() {

	nodeID := "127.0.0.1" //lets use localhost for nodeID for now

	blockchain := newBlockChain(nodeID)

	blockchain.addBlock("Teste 1")
	blockchain.addBlock("Teste 2")
	blockchain.addBlock("Teste 3")

	blockchain.print()

	return
}

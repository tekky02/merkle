package main

import (
	"merkle/merkletree"
	"merkle/socket"
)

func main() {
	mtree, _ := merkletree.NewMerkleTree("./data")
	mtree.Show()
	socket.Say()
}

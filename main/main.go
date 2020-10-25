/**
 * @ Author: tekky
 * @ Create Time: 2020-10-26 05:41:55
 * @ Modified by: tekky
 * @ Modified time: 2020-10-26 07:42:40
 * @ Description: handle data synchronization.
 */

package main

import (
	"fmt"
	"time"

	"merkle/merkletree"
)

func main() {
	mtree, _ := merkletree.NewMerkleTree("./data")
	mtree.Show()
	var c chan int
	var diff []*merkletree.Node
	for {
		diff = diff[:0]
		select {
		case <-time.After(20 * time.Second):
			fmt.Println("20 seconds passed, check...")
			newMtree, _ := merkletree.NewMerkleTree("./data")
			if mtree.Compare(newMtree, diff) {
				fmt.Println("nothing changed!")
			} else {
				fmt.Printf("file modefied!")
				diff[0].ShowContent()
			}
		case m := <-c:
			handle(m)
		}
	}
}

func handle(m int) {}

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
	for {
		select {
		case <-time.After(20 * time.Second):
			fmt.Println("20 seconds passed, check...")
			newMtree, _ := merkletree.NewMerkleTree("./data")
			if d1, d2 := mtree.Compare(newMtree); d1 == nil && d2 == nil {
				fmt.Println("nothing changed!")
			} else {
				fmt.Println("something changed")
				fmt.Println("============Old Content================")
				d1.ShowContent()
				fmt.Println("=======================================")
				fmt.Println("============New Content================")
				d2.ShowContent()
				fmt.Println("=======================================")
				mtree = newMtree
			}
		case m := <-c:
			handle(m)
		}
	}
}

func handle(m int) {}

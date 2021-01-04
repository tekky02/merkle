/**
 * @ Author: tekky
 * @ Create Time: 2020-10-26 05:47:47
 * @ Modified by: tekky
 * @ Modified time: 2020-10-26 07:42:01
 * @ Description: datastructure of Node of merkle tree.
 */

package merkletree

import (
	"crypto/sha256"
	"fmt"
)

// Node is the iner nodes of the merkle tree.
type Node struct {
	Key      int               // key of current node, used for speed up search
	Leaf     bool              // if current node is leaf.
	content  []byte            // only leaf node none nil content, which contains file content.
	Checksum [sha256.Size]byte // hash value of this node
	Parent   *Node
	Left     *Node
	Right    *Node
}

// Equals compare if node and other has the same CheckSum.
func (node *Node) Equals(other *Node) bool {
	if node.Leaf != other.Leaf {
		return false
	}
	for i, v := range node.Checksum {
		if v != other.Checksum[i] {
			return false
		}
	}
	return true
}

// ShowContent will print the content of node.
func (node *Node) ShowContent() {
	if !node.Leaf {
		panic("not a leaf!!!")
	}
	fmt.Printf("%s", node.content)
}

package merkletree

import (
	"crypto/sha256"
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

/**
 * @ Author: tekky
 * @ Create Time: 2020-10-26 05:45:39
 * @ Modified by: tekky
 * @ Modified time: 2020-10-26 07:48:40
 * @ Description: datastructure of a merkle tree.
 */

package merkletree

import (
	"crypto/sha256"
	"io/ioutil"
  "fmt"
	"path"
)

// MerkleTree is a merkle tree.
type MerkleTree struct {
	head *Node
}

// NewMerkleTree will constructs an instance of MerkleTree.
func NewMerkleTree(path string) (*MerkleTree, error) {
	files, err := getFiles(path)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{head: initNode(files, 0, len(files)-1)}, nil
}

// Show will print merkle tree to stdout.
func (mt *MerkleTree) Show() {
	fmt.Println("show checksum of all leaves...")
	print(mt.head)
}

// Pass will do nothing.
func (mt *MerkleTree) Pass() {}

func print(node *Node) {
	if node != nil {
		if node.Leaf {
			fmt.Printf("content: %s", node.content)
			fmt.Printf("checksum: %x\n", node.Checksum)
			fmt.Println("==========================================================================")
		}
		print(node.Left)
		print(node.Right)
	}
}

func getFiles(pathdir string) ([]string, error) {
	files, err := ioutil.ReadDir(pathdir)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, v := range files {
		if v.IsDir() || v.Name()[0] == '.' {
			continue
		}
		res = append(res, path.Join(pathdir, v.Name()))
	}
	return res, nil
}

func initNode(files []string, left, right int) *Node {
	if left == right {
		content, _ := ioutil.ReadFile(files[left])
		return &Node{
			Leaf:     true,
			content:  content,
			Checksum: sha256.Sum256(content),
		}
	}
	mid := (right-left)/2 + left
	lchild := initNode(files, left, mid)
	rchild := initNode(files, mid+1, right)
	current := &Node{
		Checksum: sha256.Sum256(append(lchild.Checksum[:],
			rchild.Checksum[:]...)),
		Left:  lchild,
		Right: rchild,
	}
	lchild.Parent, rchild.Parent = current, current
	return current
}

// Compare compares if mt is the same with t.
func (mt *MerkleTree) Compare(t *MerkleTree) (*Node, *Node) {
	return compare(mt.head, t.head)
}

func compare(t1, t2 *Node) (*Node, *Node) {
	if t1 != nil && t2 != nil {
		if !t1.Equals(t2) {
			if t1.Leaf {
				return t1, t2
			} else {
				td1, td2 := compare(t1.Left, t2.Left)
				if td1 != nil && td2 != nil {
					return td1, td2
				}
				return compare(t1.Right, t2.Right)
			}
		}
	}
	return nil, nil
}

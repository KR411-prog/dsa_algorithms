// https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php

package main

import "golang.org/x/exp/errors/fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert_data(d int) {
	n := &Node{
		data: d,
	}
	if t.root == nil {
		t.root = n
	} else {
		t.root.insert(d)
	}
}

func (n *Node) insert(d int) {
	if d<= n.data {
		if n.left == nil {
			n.left = &Node{
				data: d,
			}
		} else {
			n.left.insert(d)
		}
	} else {
		if n.right == nil {
			n.right = &Node {
				data: d,
			}
		} else {
			n.right.insert(d)
		}
	}


}

func preorder(n *Node) {
	if  n == nil {
		return
	} else {
		fmt.Printf("%d ", n.data)
		preorder(n.left)
		preorder(n.right)
	}
}

func postorder(n *Node) {
	if n == nil {
		return
	} else {
		postorder(n.left)
		postorder(n.right)
		fmt.Printf("%d ", n.data)
	}
}

func inorder(n *Node) {
  if n == nil {
	return
  } else {
	 inorder(n.left)
	 fmt.Printf("%d ", n.data)
	 inorder(n.right)
  }
}

func main() {
	t := &Tree{}
	t.insert_data(10)
	t.insert_data(1)
	t.insert_data(20)
	t.insert_data(3)
	t.insert_data(5)

	preorder(t.root)
	fmt.Println()
    postorder(t.root)
	fmt.Println()
	inorder(t.root)
}
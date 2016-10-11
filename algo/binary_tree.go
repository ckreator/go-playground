package main

import "fmt"

type any interface{}

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root   *Node
	sorter func(a, b int) int
}

func _rec(tree *Tree, curr *Node, v int) {
	if curr == nil {
		curr = new(Node)
		curr.val = v
		return
	}
	sorted := tree.sorter(curr.val, v)
	if sorted == 0 {
		return
	} else if sorted < 0 {
		if curr.left == nil {
			curr.left = new(Node)
			curr.left.val = v
			return
		} else {
			_rec(tree, curr.left, v)
		}
	} else {
		if curr.right == nil {
			curr.right = new(Node)
			curr.right.val = v
			return
		} else {
			_rec(tree, curr.right, v)
		}
	}
}

func (tree *Tree) addTree(val int) {
	// traverse the tree recursively

	_rec(tree, tree.root, val)
	//res := tree.sorter(tree.root.val, val)
	//return recursiveDo(tree.root, val)
}

func _printTree(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	_printTree(node.right)
	_printTree(node.left)
}

func printTree(tree *Tree) {
	if tree == nil {
		return
	}
	fmt.Print("Tree: ")
	_printTree(tree.root)
	fmt.Print("\n")
}

func intSort(a, b int) int {
	fmt.Println("a, b: ", a, b)
	return (a - b)
}

func newTree(init int) *Tree {
	tree := new(Tree)
	tree.sorter = intSort
	tree.root = new(Node)
	tree.root.val = init
	return tree
}

func main() {
	//r := Node{val: 5}
	tree := newTree(5)
	tree.addTree(10)
	tree.addTree(11)
	tree.addTree(4)
	tree.addTree(0)
	//tree.addTree(19)
	printTree(tree)
}

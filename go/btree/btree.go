package main

import (
	"fmt"
)

type tree struct {
	left, right *tree
	data        int
}

//inorder traversal
func traverse(t *tree) { // passing the root and channel
	// fmt.Println("in traverse", t.data)
	if t == nil {
		return
	} else {
		traverse(t.left)
		fmt.Println(t.data)
		traverse(t.right)
	}
}

func insert(t *tree, l *tree) *tree {
	if t == nil {
		t = l
		return t
	}
	if l.data < t.data {
		fmt.Println("Inserting into left")
		t.left = insert(t.left, l)
	} else {
		fmt.Println("Inserting into right")
		t.right = insert(t.right, l)
	}
	return t
}

func main() {
	var t = new(tree)
	k := 80
	t.data = k
	var newnode = new(tree)
	newnode.data = 20
	t = insert(t, newnode)
	var newnode1 = new(tree)
	newnode1.data = 90
	t = insert(t, newnode1)
	traverse(t)
}

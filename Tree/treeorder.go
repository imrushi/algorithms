package main

import "fmt"

type Node struct {
	num   int
	Left  *Node
	Right *Node
}

var visited []int

func preOrder(node *Node) {
	if node != nil {
		visited = append(visited, node.num)

		if node.Left != nil {
			preOrder(node.Left)
		}

		if node.Right != nil {
			preOrder(node.Right)
		}
	}
}

var invisited []int

func inOrder(node *Node) {
	if node != nil {

		if node.Left != nil {
			inOrder(node.Left)
		}

		invisited = append(invisited, node.num)

		if node.Right != nil {
			inOrder(node.Right)
		}
	}
}

var postvisited []int

func postOrder(node *Node) {
	if node != nil {

		if node.Left != nil {
			postOrder(node.Left)
		}

		if node.Right != nil {
			postOrder(node.Right)
		}

		postvisited = append(postvisited, node.num)
	}
}

func main() {
	n := &Node{1, nil, nil}
	n.Left = &Node{2, nil, nil}
	n.Right = &Node{3, nil, nil}
	n.Left.Left = &Node{4, nil, nil}
	n.Left.Right = &Node{5, nil, nil}
	n.Right.Left = &Node{6, nil, nil}
	n.Right.Right = &Node{7, nil, nil}
	n.Left.Left.Left = &Node{8, nil, nil}
	preOrder(n)
	inOrder(n)
	postOrder(n)
	fmt.Println(visited)
	fmt.Println(invisited)
	fmt.Println(postvisited)
}
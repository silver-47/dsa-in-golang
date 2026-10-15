package tree

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

type Node struct {
	Data        int
	Left, Right *Node
}

type BinaryTree struct {
	Root *Node
	Size int
}

// Helper func to find the height of the tree recursively.
func TreeHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return max(TreeHeight(node.Left), TreeHeight(node.Right)) + 1
}

// Get the height of the tree.
func (bt BinaryTree) Height() int {
	return TreeHeight(bt.Root)
}

// Insert the data into the tree using level order traversal.
func (bt *BinaryTree) Insert(data int) {
	if bt.Root == nil {
		bt.Root = &Node{Data: data}
	} else {
		q := []*Node{bt.Root}
		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			if current.Left != nil {
				q = append(q, current.Left)
			} else {
				current.Left = &Node{Data: data}
				break
			}
			if current.Right != nil {
				q = append(q, current.Right)
			} else {
				current.Right = &Node{Data: data}
				break
			}
		}
	}
	bt.Size++
}

// Helper func to find the node and parent node recursively for a given target.
func FindNodeAndParent(root *Node, target int, parent *Node) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}
	if root.Data == target {
		return root, parent
	}
	if leftTarget, leftParent := FindNodeAndParent(root.Left, target, root); leftTarget != nil {
		return leftTarget, leftParent
	}
	if rightTarget, rightParent := FindNodeAndParent(root.Right, target, root); rightTarget != nil {
		return rightTarget, rightParent
	}
	return nil, nil
}

// Find the node and parent node for the given data.
// Return (ParentNode, FoundNode)
// Parent node is nil for root node.
func (bt BinaryTree) Search(data int) (*Node, *Node) {
	return FindNodeAndParent(bt.Root, data, nil)
}

// Delete the node with the given data.
// Replace the node to be deleted with the deepest node of the tree.
// If the node is not found, return false.
// If the node is found and deleted, return true.
func (bt *BinaryTree) Delete(data int) bool {
	if bt.Root == nil {
		return false
	}
	if bt.Root.Left == nil && bt.Root.Right == nil {
		if bt.Root.Data == data {
			bt.Root = nil
			bt.Size--
			return true
		} else {
			return false
		}
	}
	type NodeAndParent struct {
		node, parent *Node
	}
	q := []NodeAndParent{{bt.Root, nil}}
	var foundNode *Node
	var current NodeAndParent
	for len(q) > 0 {
		current = q[0]
		q = q[1:]
		if current.node.Data == data {
			foundNode = current.node
		}
		if current.node.Left != nil {
			q = append(q, NodeAndParent{current.node.Left, current.node})
		}
		if current.node.Right != nil {
			q = append(q, NodeAndParent{current.node.Right, current.node})
		}
	}
	if foundNode != nil {
		foundNode.Data = current.node.Data
		if current.parent.Left == current.node {
			current.parent.Left = nil
		} else {
			current.parent.Right = nil
		}
		bt.Size--
		return true
	}
	return false
}

// Enum for traversal types.
type TraversalType int

const (
	InOrderTraversal TraversalType = iota
	PreOrderTraversal
	PostOrderTraversal
	LevelOrderTraversal
)

// Helper func to traverse the tree recursively.
func traverse(traversalType TraversalType, node *Node, res *[]int) {
	if node == nil {
		return
	}
	switch traversalType {
	case InOrderTraversal:
		traverse(InOrderTraversal, node.Left, res)
		*res = append(*res, node.Data)
		traverse(InOrderTraversal, node.Right, res)
	case PreOrderTraversal:
		*res = append(*res, node.Data)
		traverse(PreOrderTraversal, node.Left, res)
		traverse(PreOrderTraversal, node.Right, res)
	case PostOrderTraversal:
		traverse(PostOrderTraversal, node.Left, res)
		traverse(PostOrderTraversal, node.Right, res)
		*res = append(*res, node.Data)
	}
}

// Get the in-order traversal of the tree.
func (bt BinaryTree) InOrder() []int {
	result := []int{}
	traverse(InOrderTraversal, bt.Root, &result)
	return result
}

// Get the pre-order traversal of the tree.
func (bt BinaryTree) PreOrder() []int {
	result := []int{}
	traverse(PreOrderTraversal, bt.Root, &result)
	return result
}

// Get the post-order traversal of the tree.
func (bt BinaryTree) PostOrder() []int {
	result := []int{}
	traverse(PostOrderTraversal, bt.Root, &result)
	return result
}

// Get the level-order traversal of the tree.
func (bt BinaryTree) LevelOrder() [][]int {
	if bt.Root == nil {
		return [][]int{}
	}
	queue, result := []*Node{bt.Root}, [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		level := []int{}
		for range size {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Data)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// Display the tree in the given traversal type.
func (bt BinaryTree) Display(traversalType TraversalType) {
	switch traversalType {
	case InOrderTraversal:
		fmt.Print("Start -> ")
		for _, v := range bt.InOrder() {
			fmt.Printf("%d -> ", v)
		}
		fmt.Println("End")
	case PreOrderTraversal:
		fmt.Print("Start -> ")
		for _, v := range bt.PreOrder() {
			fmt.Printf("%d -> ", v)
		}
		fmt.Println("End")
	case PostOrderTraversal:
		fmt.Print("Start -> ")
		for _, v := range bt.PostOrder() {
			fmt.Printf("%d -> ", v)
		}
		fmt.Println("End")
	case LevelOrderTraversal:
		levels := bt.LevelOrder()
		for i, level := range levels {
			fmt.Printf("Level %d: ", i+1)
			for _, v := range level {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
	}
}

func SimpleBinaryTree() {
	var t BinaryTree

	for {
		menu := gocliselect.NewMenu("Binary Tree").
			AddItem("Insert", "insert").
			AddItem("Search", "search").
			AddItem("Delete", "delete").
			AddItem("Display", "display").
			AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "insert":
			var data int
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			t.Insert(data)
		case "search":
			if t.Size == 0 {
				fmt.Println("Tree is empty.")
				continue
			}
			var data int
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			if node, parent := t.Search(data); node != nil {
				fmt.Printf("Data %d (node %p) was found.\n", node.Data, node)
				if parent != nil {
					fmt.Printf("And its parent is %d (node %p).\n", parent.Data, parent)
				} else {
					fmt.Println("And it's the root node.")
				}
			} else {
				fmt.Println("Data not found.")
			}
		case "delete":
			if t.Size == 0 {
				fmt.Println("Tree is empty.")
				continue
			}
			var data int
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			if t.Delete(data) {
				fmt.Println("Data deleted.")
			} else {
				fmt.Println("Data not found.")
			}
		case "display":
			if t.Size == 0 {
				fmt.Println("Tree is empty.")
				continue
			}
			menu := gocliselect.NewMenu("Display").
				AddItem("In Order", "in").
				AddItem("Pre Order", "pre").
				AddItem("Post Order", "post").
				AddItem("Level Order", "level").
				AddItem("Back to Main Menu", "back")

			subChoice := menu.Display()

			switch subChoice {
			case "in":
				t.Display(InOrderTraversal)
			case "pre":
				t.Display(PreOrderTraversal)
			case "post":
				t.Display(PostOrderTraversal)
			case "level":
				t.Display(LevelOrderTraversal)
			case "back":
				continue
			}
		case "back":
			return
		}
	}
}

package stack

import (
	"errors"
	"fmt"

	"github.com/nexidian/gocliselect"
)

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top  *Node
	Size int
}

// Push a node at the top of the stack.
func (stack *Stack) Push(node *Node) {
	if stack.IsEmpty() {
		stack.Top = node
	} else {
		node.Next = stack.Top
		stack.Top = node
	}
	stack.Size++
}

// Pop a node from the top of the stack.
func (stack *Stack) Pop() (*Node, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}
	node := stack.Top
	stack.Top = stack.Top.Next
	stack.Size--
	return node, nil
}

// Peek at the top node of the stack.
func (stack Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty.")
	}
	return stack.Top.Data, nil
}

// Check if the stack is empty.
func (stack Stack) IsEmpty() bool {
	return stack.Size == 0
}

func SimpleStack() {
	var s = Stack{}

	for {
		menu := gocliselect.NewMenu("Stack Operations").
			AddItem("Push", "push").
			AddItem("Pop", "pop").
			AddItem("Peek", "peek").
			AddItem("Display", "display").
			AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "push":
			newNode := &Node{}
			fmt.Print("Enter the value to push: ")
			fmt.Scanln(&newNode.Data)
			s.Push(newNode)
		case "pop":
			node, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Popped value: %d\n", node.Data)
			}
		case "peek":
			data, err := s.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Top value: %d\n", data)
			}
		case "display":
			if s.IsEmpty() {
				fmt.Println("Stack is empty.")
			} else {
				fmt.Print("TOP -> ")
				for current := s.Top; current != nil; current = current.Next {
					fmt.Print(current.Data, " -> ")
				}
				fmt.Println("END")
			}
		case "back":
			return
		}
	}
}

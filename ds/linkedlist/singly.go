package linkedlist

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

type SinglyNode struct {
	Data int
	Next *SinglyNode
}

type SinglyLinkedList struct {
	Head *SinglyNode
	Size int
}

func (list *SinglyLinkedList) Display() {
	if list.Size == 0 {
		fmt.Println("List is empty.")
	} else {
		current := list.Head
		fmt.Print("HEAD -> ")
		for current != nil {
			fmt.Print(current.Data, " -> ")
			current = current.Next
		}
		fmt.Println("NULL")
	}
}

// Insert node at specific position of list.
// 1 is the first position.
// Supports negative indexing.
// -1 is the last position.
func (list *SinglyLinkedList) Insert(pos int, node *SinglyNode) {
	if pos < 0 {
		pos = list.Size + pos + 2
	}
	if pos < 1 || pos > list.Size+1 {
		fmt.Println("Invalid position.")
	} else {
		if pos == 1 {
			node.Next = list.Head
			list.Head = node
		} else {
			current := list.Head
			for i := 1; i < pos-1; i++ {
				current = current.Next
			}
			node.Next = current.Next
			current.Next = node
		}
		list.Size++
	}
}

// Delete node at specific position of list.
// 1 is the first position.
// Supports negative indexing.
// -1 is the last position.
func (list *SinglyLinkedList) Delete(pos int) {
	if pos < 0 {
		pos = list.Size + pos + 1
	}
	if pos < 1 || pos > list.Size {
		fmt.Println("Invalid position.")
	} else {
		if pos == 1 {
			list.Head = list.Head.Next
		} else {
			current := list.Head
			for i := 1; i < pos-1; i++ {
				current = current.Next
			}
			current.Next = current.Next.Next
		}
		list.Size--
	}
}

func Singly() {
	var sll = SinglyLinkedList{}

	for {
		menu := gocliselect.NewMenu("Singly Linked List Operations")

		menu.AddItem("Insert", "insert")
		menu.AddItem("Delete", "delete")
		menu.AddItem("Display", "display")
		menu.AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "display":
			sll.Display()
		case "insert":
			newNode := &SinglyNode{}
			fmt.Print("Enter the value to insert: ")
			fmt.Scanln(&newNode.Data)
			if sll.Size == 0 {
				sll.Head = newNode
				sll.Size = 1
			} else {
				submenu := gocliselect.NewMenu("Where do you want to insert the value at?")

				submenu.AddItem("At Specific Position", "position")
				submenu.AddItem("At Beginning", "beginning")
				submenu.AddItem("At End", "end")

				subChoice := submenu.Display()

				switch subChoice {
				case "beginning":
					sll.Insert(1, newNode)
				case "end":
					sll.Insert(-1, newNode)
				case "position":
					var position int
					fmt.Print("Enter the position to insert at: ")
					fmt.Scanln(&position)
					sll.Insert(position, newNode)
				}
			}
		case "delete":
			if sll.Size == 0 {
				fmt.Println("Please insert some values into the list first.")
			} else if sll.Size == 1 {
				sll.Head = nil
				sll.Size = 0
			} else {
				submenu := gocliselect.NewMenu("Where do you want to delete the value from?")

				submenu.AddItem("From Specific Position", "position")
				submenu.AddItem("From Beginning", "beginning")
				submenu.AddItem("From End", "end")

				subChoice := submenu.Display()

				switch subChoice {
				case "beginning":
					sll.Delete(1)
				case "end":
					sll.Delete(-1)
				case "position":
					var position int
					fmt.Print("Enter the position to delete from: ")
					fmt.Scanln(&position)
					sll.Delete(position)
				}
			}
		case "back":
			return
		}
	}
}

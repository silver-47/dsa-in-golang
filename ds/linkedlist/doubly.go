package linkedlist

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

type DoublyNode struct {
	Data       int
	Next, Prev *DoublyNode
}

type DoublyLinkedList struct {
	Head, Tail *DoublyNode
	Size       int
}

// Display all nodes of the list.
// reverse: false to display the list in normal order.
// reverse: true to display the list in reverse order.
func (list DoublyLinkedList) Display(reverse bool) {
	if list.Size == 0 {
		fmt.Println("List is empty.")
	} else {
		current := list.Head
		if reverse {
			current = list.Tail
		}
		fmt.Print("NULL <-> ")
		for current != nil {
			fmt.Print(current.Data, " <-> ")
			if reverse {
				current = current.Prev
			} else {
				current = current.Next
			}
		}
		fmt.Println("NULL")
		fmt.Println("         ꜛ")
		if reverse {
			fmt.Println("        TAIL")
		} else {
			fmt.Println("        HEAD")
		}
	}
}

// Insert node at given position of the list.
// 1 is the first position.
// Supports negative indexing.
// -1 is the last position.
func (list *DoublyLinkedList) Insert(pos int, node *DoublyNode) {
	if pos < 0 {
		pos = list.Size + pos + 2
	}
	if pos < 1 || pos > list.Size+1 {
		fmt.Println("Invalid position.")
	} else {
		if pos == 1 {
			node.Next = list.Head
			node.Prev = nil
			list.Head.Prev = node
			list.Head = node
		} else if pos == list.Size+1 {
			node.Prev = list.Tail
			node.Next = nil
			list.Tail.Next = node
			list.Tail = node
		} else if pos <= list.Size/2 {
			current := list.Head
			for i := 1; i < pos-1; i++ {
				current = current.Next
			}
			node.Next = current.Next
			node.Prev = current
			current.Next.Prev = node
			current.Next = node
		} else {
			current := list.Tail
			for i := list.Size; i > pos; i-- {
				current = current.Prev
			}
			node.Prev = current.Prev
			node.Next = current
			current.Prev.Next = node
			current.Prev = node
		}
		list.Size++
	}
}

// Delete node at specific position of the list.
// 1 is the first position.
// Supports negative indexing.
// -1 is the last position.
func (list *DoublyLinkedList) Delete(pos int) {
	if pos < 0 {
		pos = list.Size + pos + 1
	}
	if pos < 1 || pos > list.Size {
		fmt.Println("Invalid position.")
	} else {
		if pos == 1 {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		} else if pos == list.Size {
			list.Tail.Prev.Next = nil
			list.Tail = list.Tail.Prev
		} else if pos <= list.Size/2 {
			current := list.Head
			for i := 1; i < pos-1; i++ {
				current = current.Next
			}
			current.Next = current.Next.Next
			current.Next.Prev = current
		} else {
			current := list.Tail
			for i := list.Size - 1; i > pos; i-- {
				current = current.Prev
			}
			current.Prev = current.Prev.Prev
			current.Prev.Next = current
		}
		list.Size--
	}
}

func Doubly() {
	var dll = DoublyLinkedList{}

	for {
		menu := gocliselect.NewMenu("Doubly Linked List Operations")

		menu.AddItem("Insert", "insert")
		menu.AddItem("Delete", "delete")
		menu.AddItem("Display", "display")
		menu.AddItem("Back to Main Menu", "back")

		subChoice := menu.Display()

		switch subChoice {
		case "display":
			submenu := gocliselect.NewMenu("Direction of Display?")
			submenu.AddItem("Normal", "normal")
			submenu.AddItem("Reverse", "reverse")
			subChoice := submenu.Display()
			dll.Display(subChoice == "reverse")
		case "insert":
			newNode := &DoublyNode{}
			fmt.Print("Enter the value to insert: ")
			fmt.Scanln(&newNode.Data)
			if dll.Size == 0 {
				dll.Head = newNode
				dll.Tail = newNode
				dll.Size = 1
			} else {
				submenu := gocliselect.NewMenu("Where do you want to insert the value at?")
				submenu.AddItem("At Specific Position", "position")
				submenu.AddItem("At Beginning", "beginning")
				submenu.AddItem("At End", "end")

				subChoice := submenu.Display()

				switch subChoice {
				case "beginning":
					dll.Insert(1, newNode)
				case "end":
					dll.Insert(-1, newNode)
				case "position":
					var position int
					fmt.Print("Enter the position to insert at: ")
					fmt.Scanln(&position)
					dll.Insert(position, newNode)
				}
			}
		case "delete":
			if dll.Size == 0 {
				fmt.Println("Please insert some values into the list first.")
			} else if dll.Size == 1 {
				dll.Head = nil
				dll.Tail = nil
				dll.Size = 0
			} else {
				submenu := gocliselect.NewMenu("Where do you want to delete the value from?")
				submenu.AddItem("From Specific Position", "position")
				submenu.AddItem("From Beginning", "beginning")
				submenu.AddItem("From End", "end")

				subChoice := submenu.Display()

				switch subChoice {
				case "beginning":
					dll.Delete(1)
				case "end":
					dll.Delete(-1)
				case "position":
					var position int
					fmt.Print("Enter the position to delete from: ")
					fmt.Scanln(&position)
					dll.Delete(position)
				}
			}
		case "back":
			return
		}
	}
}

package main

import (
	"fmt"

	"github.com/nexidian/gocliselect"
	"github.com/silver-47/dsa-in-golang/ds/array"
	"github.com/silver-47/dsa-in-golang/ds/linkedlist"
	"github.com/silver-47/dsa-in-golang/ds/stack"
)

func main() {
	for {
		menu := gocliselect.NewMenu("Data Structures")

		menu.AddItem("Array", "array")
		menu.AddItem("Linked List", "linkedlist")
		menu.AddItem("Stack", "stack")
		menu.AddItem("Queue", "queue")
		menu.AddItem("Tree", "tree")
		menu.AddItem("Graph", "graph")
		menu.AddItem("Heap", "heap")
		menu.AddItem("Hash Table", "hashtable")
		menu.AddItem("Exit", "exit")

		choice := menu.Display()

		switch choice {
		case "array":
			array.Array()
		case "linkedlist":
			submenu := gocliselect.NewMenu("Linked List")

			submenu.AddItem("Singly Linked List", "singly")
			submenu.AddItem("Doubly Linked List", "doubly")
			submenu.AddItem("Back to Main Menu", "back")

			subChoice := submenu.Display()

			switch subChoice {
			case "singly":
				linkedlist.Singly()
			case "doubly":
				linkedlist.Doubly()
			case "back":
				continue
			}
		case "stack":
			stack.SimpleStack()
		case "exit":
			fmt.Println("Thank you for using this program.")
			return
		}
	}
}

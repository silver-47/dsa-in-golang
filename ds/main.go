package main

import (
	"fmt"

	"github.com/nexidian/gocliselect"
	"github.com/silver-47/dsa-in-golang/ds/array"
	"github.com/silver-47/dsa-in-golang/ds/linkedlist"
	"github.com/silver-47/dsa-in-golang/ds/queue"
	"github.com/silver-47/dsa-in-golang/ds/stack"
)

func main() {
	for {
		menu := gocliselect.NewMenu("Data Structures").
			AddItem("Array", "array").
			AddItem("Linked List", "linkedlist").
			AddItem("Stack", "stack").
			AddItem("Queue", "queue").
			AddItem("Tree", "tree").
			AddItem("Graph", "graph").
			AddItem("Heap", "heap").
			AddItem("Hash Table", "hashtable").
			AddItem("Exit", "exit")

		choice := menu.Display()

		switch choice {
		case "array":
			array.Array()
		case "linkedlist":
			submenu := gocliselect.NewMenu("Linked List").
				AddItem("Singly Linked List", "singly").
				AddItem("Doubly Linked List", "doubly").
				AddItem("Back to Main Menu", "back")

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
		case "queue":
			submenu := gocliselect.NewMenu("Queue").
				AddItem("Simple Queue", "simple").
				AddItem("Circular Queue", "circular").
				// submenu.AddItem("Priority Queue", "priority")
				AddItem("Back to Main Menu", "back")

			subChoice := submenu.Display()

			switch subChoice {
			case "simple":
				queue.SimpleQueue()
			case "circular":
				queue.CircularQueue()
			case "priority":
				queue.PriorityQueue()
			case "back":
				continue
			}
		case "exit":
			fmt.Println("Thank you for using this program.")
			return
		}
	}
}

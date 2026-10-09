package main

import (
	"fmt"

	"github.com/nexidian/gocliselect"
	"github.com/silver-47/dsa-in-golang/ds/array"
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
		case "exit":
			fmt.Println("Thank you for using this program.")
			return
		}
	}
}

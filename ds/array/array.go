package array

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

func Array() {
	var size int

	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)

	arr := make([]int, size)

	for {
		menu := gocliselect.NewMenu("Array Operations")

		menu.AddItem("Insert", "insert")
		menu.AddItem("Delete", "delete")
		menu.AddItem("Display", "display")
		menu.AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "insert":
			element := 0
			fmt.Print("Enter the element: ")
			fmt.Scan(&element)
			index := 0
			fmt.Print("Enter the index: ")
			fmt.Scan(&index)

			if index < 0 || index >= size {
				fmt.Println("Invalid index")
				return
			}

			arr[index] = element
		case "delete":
			index := 0
			fmt.Print("Enter the index: ")
			fmt.Scan(&index)

			if index < 0 || index >= size {
				fmt.Println("Invalid index")
				return
			}

			arr[index] = 0
		case "display":
			fmt.Println("Array is:", arr)
		case "back":
			return
		}
	}
}

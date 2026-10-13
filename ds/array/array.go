package array

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

func Array() {
	var size int

	for {
		fmt.Print("Enter the size of the array: ")
		fmt.Scan(&size)
		if size < 1 {
			fmt.Println("Invalid size")
		} else {
			break
		}
	}

	arr := make([]int, size)

	for {
		menu := gocliselect.NewMenu("Array Operations").
			AddItem("Insert", "insert").
			AddItem("Delete", "delete").
			AddItem("Display", "display").
			AddItem("Back to Main Menu", "back")

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
				continue
			}

			arr[index] = element
		case "delete":
			index := 0
			fmt.Print("Enter the index: ")
			fmt.Scan(&index)

			if index < 0 || index >= size {
				fmt.Println("Invalid index")
				continue
			}

			arr[index] = 0
		case "display":
			fmt.Println("Array is:", arr)
		case "back":
			return
		}
	}
}

package queue

import (
	"fmt"

	"github.com/nexidian/gocliselect"
)

func CircularQueue() {
	var capacity int

	for {
		fmt.Print("Enter the capacity of the circular queue: ")
		fmt.Scan(&capacity)
		if capacity < 1 {
			fmt.Println("Invalid capacity")
		} else {
			break
		}
	}

	q, front, rear := make([]int, capacity), -1, -1

	for {
		menu := gocliselect.NewMenu("Circular Queue")

		menu.AddItem("Enqueue", "enqueue")
		menu.AddItem("Dequeue", "dequeue")
		menu.AddItem("Peek", "peek")
		menu.AddItem("Display", "display")
		menu.AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "enqueue":
			if (rear+1)%capacity == front {
				fmt.Println("Queue Overflow!")
			} else {
				var element int
				fmt.Print("Enter the element to enqueue: ")
				fmt.Scan(&element)

				rear = (rear + 1) % capacity
				q[rear] = element
				if front == -1 {
					front = 0
				}
			}
		case "dequeue":
			if front == -1 || rear == -1 {
				fmt.Println("Queue Underflow!")
			} else {
				fmt.Printf("Dequeued value: %d\n", q[front])
				if front == rear {
					front = -1
					rear = -1
				} else {
					front = (front + 1) % capacity
				}
			}
		case "peek":
			if front == -1 || rear == -1 {
				fmt.Println("Queue is empty.")
			} else {
				fmt.Printf("Value at the front: %d\n", q[front])
			}
		case "display":
			if front == -1 || rear == -1 {
				fmt.Println("Queue is empty.")
			} else {
				fmt.Print("FRONT <- ")
				for i := front; i != rear; i = (i + 1) % capacity {
					fmt.Printf("%d <- ", q[i])
				}
				fmt.Printf("%d <- REAR\n", q[rear])
			}
		case "back":
			return
		}
	}
}

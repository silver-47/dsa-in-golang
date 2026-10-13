package queue

import (
	"errors"
	"fmt"

	"github.com/nexidian/gocliselect"
)

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Front, Rear *Node
	Size        int
}

func (q *Queue) Enqueue(node *Node) {
	if q.IsEmpty() {
		q.Front = node
		q.Rear = node
	} else {
		q.Rear.Next = node
		q.Rear = node
	}
	q.Size++
}

func (q *Queue) Dequeue() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}
	node := q.Front
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
	q.Size--
	return node, nil
}

func (q Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty.")
	}
	return q.Front.Data, nil
}

func (q Queue) IsEmpty() bool {
	return q.Size == 0
}

func SimpleQueue() {
	var queue Queue

	for {
		menu := gocliselect.NewMenu("Queue").
			AddItem("Enqueue", "enqueue").
			AddItem("Dequeue", "dequeue").
			AddItem("Peek", "peek").
			AddItem("Display", "display").
			AddItem("Back to Main Menu", "back")

		choice := menu.Display()

		switch choice {
		case "enqueue":
			newNode := &Node{}
			fmt.Print("Enter the value to enqueue: ")
			fmt.Scanln(&newNode.Data)
			queue.Enqueue(newNode)
		case "dequeue":
			node, err := queue.Dequeue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Dequeued value: %d\n", node.Data)
			}
		case "peek":
			value, err := queue.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Value at the front: %d\n", value)
			}
		case "display":
			if queue.IsEmpty() {
				fmt.Println("Queue is empty.")
			} else {
				fmt.Print("FRONT <- ")
				for current := queue.Front; current != nil; current = current.Next {
					fmt.Print(current.Data, " <- ")
				}
				fmt.Println("REAR")
			}
		case "back":
			return
		}
	}
}

package main

import "fmt"
import sc "strconv"

func main() {
	queue := Queue{}

	queue.add(5)
	queue.add(3)
	queue.add(2)

	fmt.Println(queue.count)

	fmt.Println(queue.exists(3))
	fmt.Println(queue.exists(12))

	// an iterator would be fine :(
	for elm, error := queue.dequeue(); !error; elm, error = queue.dequeue() {
		fmt.Println("Got element from the queue: " + sc.Itoa(elm.(int)))
		fmt.Println("Remaining elements: " + sc.Itoa(queue.count))
	}
}

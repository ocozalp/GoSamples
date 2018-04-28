package main

import "fmt"
import sc "strconv"

type Node struct {
    val interface{}
    next *Node
}

type Queue struct {
    head *Node
    tail *Node
    count int
}

// This is unexpected
func (queue *Queue) add(val interface{}) {
    newNode := Node{val:val, next:nil}
    
    if queue.head == nil {
        queue.head = &newNode
        queue.tail = &newNode
    } else {
        queue.tail.next = &newNode
        queue.tail = &newNode
    }

    queue.count++ 
}

func (queue *Queue) exists(val interface{}) bool {
    current := queue.head

    for current != nil {
        if current.val == val {
            return true
        }

        current = current.next
    }

    return false
}

func (queue *Queue) dequeue() (result interface{}, error bool) {
    currentHeadNode := queue.head

    if currentHeadNode == nil {
        error = true
        return
    }
    
    nextNode := queue.head.next

    queue.head = nextNode
    if nextNode == nil {
        queue.tail = nil
    }

    queue.count--

    return currentHeadNode.val, false
}

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

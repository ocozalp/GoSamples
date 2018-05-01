package main

type Node struct {
	val  interface{}
	next *Node
}

type Queue struct {
	head  *Node
	tail  *Node
	count int
}

// This is unexpected
func (queue *Queue) add(val interface{}) {
	newNode := Node{val: val, next: nil}

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

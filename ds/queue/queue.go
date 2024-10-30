package queue

type Queue struct {
	rear   *node
	front  *node
	length int
}

type node struct {
	value any
	next  *node
}

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (queue *Queue) Enqueue(val any) {
	n := &node{val, nil}
	if queue.length == 0 {
		queue.front = n
	} else {
		queue.rear.next = n
	}
	queue.rear = n
	queue.length++
}

func (queue *Queue) Dequeue() any {
	if queue.length == 0 {
		return nil
	}
	n := queue.front
	queue.front = n.next
	queue.length--
	return n.value
}

func (queue *Queue) Peek() any {
	if queue.length == 0 {
		return nil
	}
	return queue.front.value
}

func (queue *Queue) Len() int {
	return queue.length
}

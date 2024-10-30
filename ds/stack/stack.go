package stack

import "iter"

type Stack struct {
	top    *node
	length int
}

type node struct {
	value any
	prev  *node
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Push(val any) {
	n := &node{val, stack.top}
	stack.top = n
	stack.length++
}

func (stack *Stack) Pop() any {
	if stack.length == 0 {
		return nil
	}
	node := stack.top
	stack.top = node.prev
	stack.length--
	return node.value
}

func (stack *Stack) Len() int {
	return stack.length
}

func (stack *Stack) Peek() any {
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
}

func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) Items() iter.Seq2[int, any] {
	return func(yield func(int, any) bool) {
		n := stack.top
		for i := 0; i < stack.length; i++ {
			if !yield(i, n.value) {
				return
			}
			n = n.prev
		}
	}
}

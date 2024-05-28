package main

import (
  "fmt"
  "github.com/sanjay-sol/learning-go"
)

// structs

type ListNode struct {
	val  int64
	next *ListNode
}

func (l *ListNode) Add(val int64) {
	if l.next == nil {
		l.next = &ListNode{val: val}
	} else {
		l.next.Add(val)
	}
}

func (l *ListNode) Print() {
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

func (l *ListNode) Reverse() {
	var prev *ListNode
	curr := l
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l = prev

}




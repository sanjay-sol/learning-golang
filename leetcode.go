package main

// structs 

types ListNode struct {
  val int 
  next *ListNode
}


func (l *ListNode) add(val int) {
  if l.next == nil {
    l.next = &ListNode{val: val}
  } else {
    l.next.add(val)
  }
}

func (l *ListNode) print() {
  for l != nil {
    fmt.Println(l.val)
    l = l.next
  }
}




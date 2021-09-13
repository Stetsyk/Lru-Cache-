package list

import "fmt"

// double linked list implememntation
type List struct {
	Key        int
	Value      int
	prev, next *List
}

func (head *List) Append(key int, value int) (retHead *List, node *List) {
	nxt := &List{Key: key, Value: value}
	if head == nil {
		head = nxt
	} else if head.prev == nil {
		// two elements each pointing on each other
		head.next = nxt
		head.prev = nxt
		nxt.next = head
		nxt.prev = head
	} else {
		last := head.prev
		head.prev = nxt // last value now is nxt
		nxt.prev = last // nxt previous value is now previous last value
		last.next = nxt // last previous value next is nxt value
		nxt.next = head // nxt next value is head
	}
	retHead = head
	node = head
	if head.prev != nil {
		node = head.prev
	}
	return retHead, node
}

func (head *List) Erase(node *List) (retHead *List) {
	if head == nil {
		return
	} else if head.next == nil {
		head = nil
	} else if *head.next.next == *head {
		if *head == *node {
			head = head.next
		}
		head.next = nil
		head.prev = nil
	} else {
		if *node == *head {
			head = head.next
		}
		prev := node.prev
		nxt := node.next
		prev.next = nxt
		nxt.prev = prev
	}
	retHead = head
	return retHead
}

func (head *List) Print() {
	headCopy := head
	fmt.Println("Printin list")
	for {
		if head == nil {
			break
		}
		fmt.Println(head)
		head = head.next
		if head == nil {
			break
		}
		if *head == *headCopy {
			break
		}
	}
	fmt.Println("End printing")
}

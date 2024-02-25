package LinkedList

import "fmt"

type ListNode struct {
	head     *Node
	capacity int
}

func (l *ListNode) Info() {
	var head = l.head
	for head != nil {
		fmt.Printf("%#v\n", head.val)
		head = head.next
	}
}

func (l *ListNode) add(val int) {
	newNode := &Node{val: val, next: nil}
	if l.capacity == 0 {
		l.head = newNode
	} else {
		node := l.head
		for node.next != nil {
			node = node.next
		}
		node.next = newNode
	}
	l.capacity++
}

func CreatList(list []int) *ListNode {
	var ll = &ListNode{}
	for i := 0; i < len(list); i++ {
		ll.add(list[i])
	}
	return ll
}

func DeleteDuplicates(node *ListNode) *ListNode {
	var cur = node.head
	var head = node
	for cur != nil && cur.next != nil {
		if cur.val == cur.next.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}

	}
	return head
}

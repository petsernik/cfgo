package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(x int) *ListNode {
	l.Next = &ListNode{x, nil}
	return l.Next
}

func (l *ListNode) delNext() {
	if l.Next == nil {
		return
	}
	l.Next = l.Next.Next
}

// 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	init := head
	for i := 0; i < left-2; i++ {
		head = head.Next
	}
	var prev_start *ListNode
	if left > 1 {
		prev_start = head
		head = head.Next
	}
	var prev, cur, next *ListNode = nil, head, head.Next
	end := head
	for i := left; i <= right; i++ {
		cur.Next = prev
		if next != nil {
			prev, cur, next = cur, next, next.Next
		} else {
			prev, cur = cur, next
		}
	}
	end.Next = cur
	if left > 1 {
		prev_start.Next = prev
		return init
	} else {
		return prev
	}
}

func reverseList(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	var prev, cur, next *ListNode = nil, h, h.Next
	for ; ; prev, cur, next = cur, next, next.Next {
		cur.Next = prev
		if next == nil {
			break
		}
	}
	return cur
}

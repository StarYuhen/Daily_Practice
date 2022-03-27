package main

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建两个链表暂存数据
	lown := head
	lows := head
	for i := 0; i < n; n++ {
		lown = lown.Next
	}

	if lown == nil {
		// 返回下一个--int类型的默认值是0
		return lown.Next
	}

	for head.Next != nil {
		lown = lown.Next
		lows = lows.Next
	}

	lows.Next = lows.Next.Next

	return head

}

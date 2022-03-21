package main

// 删除链表的节点

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// 就把指定链表数据给替换就完事了
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

package main

// import "fmt"

// type ListNode struct {
// 	val int
// 	Next *ListNode
// }

// func constructLD(val int) *ListNode {
// 	return &ListNode{val: val, Next: nil}
// }

func middleNode(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	slowPtr := head
	fastPtr := head.Next

	for fastPtr != nil {

		slowPtr = slowPtr.Next

		if fastPtr.Next == nil {
			fastPtr = nil
		} else {
        	fastPtr = fastPtr.Next.Next
		}
	}

	return slowPtr
}

// func main() {
// 	beg := constructLD(1)
// 	beg.Next = constructLD(2)
// 	beg.Next.Next = constructLD(3)
// 	beg.Next.Next.Next = constructLD(4)
// 	beg.Next.Next.Next.Next = constructLD(5)
// 	beg.Next.Next.Next.Next.Next = constructLD(6)
// 	fmt.Println(middleNode(beg))
// }
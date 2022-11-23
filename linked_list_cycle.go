package main

// import "fmt"

type ListNode struct {
	val int
	Next *ListNode
}

func constructLD(val int) *ListNode {
	return &ListNode{val: val, Next: nil}
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slowPointer := head
    fastPointer := head.Next

    for slowPointer != nil && fastPointer != nil {
        if slowPointer == fastPointer {
            return true
        }

        if slowPointer.Next == nil || fastPointer.Next == nil {
            return false
        }

        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next
    }

    return false
}

// func main() {
// 	beg := constructLD(3)
// 	beg.Next = constructLD(2)
// 	beg.Next.Next = constructLD(0)
// 	beg.Next.Next.Next = constructLD(-4)
// 	beg.Next.Next.Next.Next = beg.Next
// 	fmt.Println(hasCycle(beg))
// }
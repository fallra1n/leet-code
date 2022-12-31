package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	res := &ListNode{0, nil}
	resCopy := res
	flag := false
	inMind := 0

	for l1 != nil && l2 != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		res.Val = (inMind + l1.Val + l2.Val) % 10
		inMind = (inMind + l1.Val + l2.Val) / 10

		l1 = l1.Next
		l2 = l2.Next

		flag = true
	}

	for l1 != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		res.Val = (inMind + l1.Val) % 10
		inMind = (inMind + l1.Val) / 10

		l1 = l1.Next
		flag = true
	}

	for l2 != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		res.Val = (inMind + l2.Val) % 10
		inMind = (inMind + l2.Val) / 10

		l2 = l2.Next
		flag = true
	}

	if inMind != 0 {
		res.Next = &ListNode{inMind, nil}
	}

	return resCopy
}

package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoList(lhs *ListNode, rhs *ListNode) *ListNode {
	if lhs == nil {
		return rhs
	}

	if rhs == nil {
		return lhs
	}

	res := &ListNode{0, nil}
	resCopy := res

	flag := false

	for lhs != nil && rhs != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		if lhs.Val < rhs.Val {
			res.Val = lhs.Val
			lhs = lhs.Next
		} else {
			res.Val = rhs.Val
			rhs = rhs.Next
		}

		flag = true
	}

	for lhs != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		res.Val = lhs.Val
		lhs = lhs.Next

		flag = true
	}

	for rhs != nil {
		if flag {
			res.Next = &ListNode{0, nil}
			res = res.Next
		}

		res.Val = rhs.Val
		rhs = rhs.Next

		flag = true
	}

	return resCopy
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	temp := lists
	for len(temp) > 1 {
		newLists := make([]*ListNode, 0)

		for i := 0; i < len(temp)-1; i += 2 {
			newLists = append(newLists, mergeTwoList(temp[i], temp[i+1]))
		}

		if len(temp)%2 == 1 {
			newLists = append(newLists, temp[len(temp)-1])
		}

		temp = newLists
	}

	return temp[0]
}

//func main() {
//	lists := make([]*ListNode, 3)
//	lists[0] = &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
//	lists[1] = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
//	lists[2] = &ListNode{2, &ListNode{6, nil}}
//
//	mergeKLists(lists)
//}

package task2

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstResultNode, prevResultNode *ListNode
	leftNode := l1
	rightNode := l2
	currentOverflow := 0

	// If both are nil - result is ready, return first node
	for leftNode != nil || rightNode != nil {

		// if one of numbers is nil it means that there are no more digits in
		// this number, so we add not significant 0 that won't affect the result
		if leftNode == nil {
			leftNode = &ListNode{Val: 0, Next: nil}
		}
		if rightNode == nil {
			rightNode = &ListNode{Val: 0, Next: nil}
		}
		// Sum numbers and overflow, if we have overflow save it for next
		// iteration
		currentSum := leftNode.Val + rightNode.Val + currentOverflow
		currentOverflow = 0
		if currentSum > 9 {
			currentOverflow = 1
			currentSum -= 10
		}
		leftNode = leftNode.Next
		rightNode = rightNode.Next
		if firstResultNode == nil {
			firstResultNode = &ListNode{Val: currentSum, Next: nil}
			prevResultNode = firstResultNode
			continue
		}
		prevResultNode.Next = &ListNode{Val: currentSum, Next: nil}
		prevResultNode = prevResultNode.Next
	}
	if currentOverflow > 0 {
		prevResultNode.Next = &ListNode{Val: currentOverflow, Next: nil}
	}
	return firstResultNode
}

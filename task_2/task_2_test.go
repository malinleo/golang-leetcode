package task2

import "testing"


func checkOutput(t *testing.T, result *ListNode, expected []int) {
	idx := 0
	for result != nil {
		if result.Val != expected[idx] {
			t.Fatalf("Error, expected %d, got %d", expected[idx], result.Val)
		}
		result = result.Next
		idx++
	}
}

func TestAddTwoNumbers1(t *testing.T) {
	l1_7 := ListNode{Val: 9, Next: nil}
	l1_6 := ListNode{Val: 9, Next: &l1_7}
	l1_5 := ListNode{Val: 9, Next: &l1_6}
	l1_4 := ListNode{Val: 9, Next: &l1_5}
	l1_3 := ListNode{Val: 9, Next: &l1_4}
	l1_2 := ListNode{Val: 9, Next: &l1_3}
	l1 := ListNode{Val: 9, Next: &l1_2}

	l2_4 := ListNode{Val: 9, Next: nil}
	l2_3 := ListNode{Val: 9, Next: &l2_4}
	l2_2 := ListNode{Val: 9, Next: &l2_3}
	l2 := ListNode{Val: 9, Next: &l2_2}
	checkOutput(t, addTwoNumbers(&l1, &l2), []int{8, 9, 9, 9, 0, 0, 0, 1})
}

func TestAddTwoNumbers2(t *testing.T) {
	l1_3 := ListNode{Val: 3, Next: nil}
	l1_2 := ListNode{Val: 4, Next: &l1_3}
	l1 := ListNode{Val: 2, Next: &l1_2}

	l2_3 := ListNode{Val: 4, Next: nil}
	l2_2 := ListNode{Val: 6, Next: &l2_3}
	l2 := ListNode{Val: 5, Next: &l2_2}
	checkOutput(t, addTwoNumbers(&l1, &l2), []int{7, 0, 8})
}
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// 时间复杂度o(n)
// 空间复杂度o(n)
func reversePrint(head *ListNode) []int {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	if len(arr) == 0 {
		return nil
	}

	length := len(arr)
	for i := 0; i < length / 2; i++ {
		temp := arr[i]
		arr[i] = arr[length - i - 1]
		arr[length - i - 1] = temp
	}

	return arr
}

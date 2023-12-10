// https://leetcode.com/problems/palindrome-linked-list

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Followed The book :)
 //Algorithm used is, Get the middle of linked list, Reverse the second half of linked list, compare the first half and second half
 import "fmt"
 func  reverseLinkedlist(head *ListNode) *ListNode {
	 var prev,temp *ListNode
	 curr := head

	 for curr != nil {
		 temp = curr.Next
		 curr.Next = prev
		 prev = curr
		 curr = temp
	 }
	 return prev
 }

 func isPalindrome(head *ListNode) bool {
	 //find length of the linked list
	 len := 0
	 curr := head
	 for curr != nil {
		 len += 1
		 curr = curr.Next
	 }

	 curr = head
	 count := 0
	 for count < len/2 {
		 curr = curr.Next
		 count += 1
	 }
	 rev_head := reverseLinkedlist(curr)
	 fmt.Println(rev_head)

	 curr = head

	 for rev_head != nil {
		 if rev_head.Val != curr.Val {
			 return false

		 }
		rev_head = rev_head.Next
		curr = curr.Next
	 }
	 return true
 }
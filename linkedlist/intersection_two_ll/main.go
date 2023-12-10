// https://leetcode.com/problems/intersection-of-two-linked-lists/description/
// implements intersection of two linked lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 import (
	"math"
	"fmt"
	)
	func getIntersectionNode(headA, headB *ListNode) *ListNode {
	   currA := headA
	   currB := headB
	   lenA,lenB := 0,0
	   for currA != nil {
		   lenA += 1
		   currA = currA.Next
	   }
		for currB != nil {
			lenB += 1
			currB = currB.Next
		}

		len := int(math.Abs(float64(lenA-lenB)))

		fmt.Println(lenA,lenB,len)

		if lenA > lenB {
			currA = headA
			currB = headB
			count := 0
			for count < len {
				currA = currA.Next
				count += 1
			}
		} else {
			currB = headB
			currA = headA
			count := 0
			for count < len {
				currB = currB.Next
				count += 1
			}
		}

		fmt.Println(currA.Val, currB.Val)

		if currA == currB {
			return currA
		}
		for currA.Next != currB.Next {
			currA = currA.Next
			currB = currB.Next
		}
		return currA.Next

	}
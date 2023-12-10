//Is LinkedLint a circular loop
//https://leetcode.com/problems/linked-list-cycle

if head == nil || head.Next == nil {
	return false
}

slow_ptr := head
fast_ptr := head.Next

for fast_ptr != slow_ptr && fast_ptr.Next != nil && fast_ptr.Next.Next != nil   {
	slow_ptr = slow_ptr.Next
	fast_ptr = fast_ptr.Next.Next

}
return slow_ptr == fast_ptr
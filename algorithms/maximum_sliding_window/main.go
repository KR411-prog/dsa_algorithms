// O(n2) time complexity
// https://leetcode.com/problems/sliding-window-maximum/
// func maxSlidingWindow(nums []int, k int) []int {
//     max := []int{}

//     for i:=0;i<len(nums)-k+1;i++ {
//         max_window := -33000
//         for j:=i;j<i+k;j++ {
//             if max_window < nums[j] {
//                 max_window = nums[j]
//             }
//         }
//             max = append(max,max_window)

//     }
//     return max
// }

//O(n) Time complexity


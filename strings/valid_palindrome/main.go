// https://leetcode.com/problems/valid-palindrome
import (
	"strings"
	"fmt"
	)

	func isValid(c byte) bool {
		return c >= 'a' && c <= 'z' || c>='0' && c<='9'
	}

	func isPalindrome(s string) bool {

		i := 0
		j := len(s)-1
		s = strings.ToLower(s)


		for i < j {


			for i<j && !isValid(s[i])  {
				i++
			}

			for i<j &&  !isValid(s[j]) {
				j--
			}


			if s[i] != s[j] {
				return false
			}

			i++
			j--
		}

		return true
	}
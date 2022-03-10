package event

import (
	"fmt"
)

func Run() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x > 0 {
		rem := x % 10
		result := 0
		result = result*10 + rem
		x = x / 10
		if x == result {
			return true
		}
	}
	return false
}

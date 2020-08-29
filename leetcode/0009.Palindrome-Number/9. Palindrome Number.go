package leetcode

// import "strconv"

func isPalindrome(x int) bool {
	y := x
	z := 0
	if x < 0 {
			return false
	}
	if x < 10 {
			return true
	}
	for y > 0 {
		i := y%10
		y = y/10;
		z = z*10 + i
	}
	if (z == x) {
		return true
	}
	return false
}

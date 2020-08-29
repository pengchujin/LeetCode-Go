package leetcode

// import (
// 	"fmt"
// )

func maxArea(height []int) int {
	max := 0
	l := 0
	for i := 0; i < (len(height) -1); i++ {
		for k := i+1; k < len(height); k++ {
			 l = height[k]
			 if height[i] <= height[k] {
				 l = height[i]
			 }
			 if l * (k-i) > max {
				 max = l * (k-i)
			 }
		}
	}
	return max
}

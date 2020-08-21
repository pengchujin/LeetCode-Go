package leetcode
import (
	"fmt"
)

func reverse7(x int) int {
	// for i := 100000; i < 2147483648; i++ {
	// 	max :=0
	// 	for z :=i; z > 0; z = z/10 {
	// 		max = (max)*10 + (z%10)
	// 	}
	// 	if max > 2147483648 {
	// 		fmt.Print("max:: ",i, max)
	// 		break
	// 	}
	// }
	res := 0
	if x > 2147483648 || x < -2147483648 {
		return res
	}
	if x > 0 {
		for z :=x; z > 0; z = z/10 {
			fmt.Println(z%10)
			res = (res)*10 + (z%10)
		}
	}
	if x < 0 {
		for z := -x; z > 0; z = z/10 {
			res = (res)*10 + (z%10)
		}
		res = -res
	}
	if res > 2147483648 || res < -2147483648 {
		return 0
	}
	return res
}

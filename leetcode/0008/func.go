package leetcode
import (
	"fmt"
)

func myAtoi(str string) int {
	key := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "0": 0}
	start := -1
	lens := 0
    tim := 0
	res := 0
	for i := 0; i < len(str); i++ {
        if (string(str[i]) == "-" || string(str[i]) == "+") && start == -1 {
            tim++
        }
         if string(str[i]) == " " && tim == 1 && start == -1{
            return 0
        }
        if _, ok := key[string(str[i])]; !ok && string(str[i]) != "-" && string(str[i]) != "+" && string(str[i]) != " " && start == -1 || tim ==2 {
		    return 0
	    }
		if val, ok := key[string(str[i])]; ok {
			start = i
			lens++
			res = (res * 10) + val
		} else if start != -1 {
			break
		}
	}
	if start == -1 {
		res =  0
	}
	fmt.Println(res)
    if (start - lens + 1) != 0 && start != -1 {
		if string(str[start - lens]) == "-" {
			res = -res
            if res < -2147483648 || res > 2147483647 {
		    return -2147483648
            } else {
                return res
            }
	}
	}
    if res > 2147483647 || res < -2147483647  {
		return 2147483647
	}
	return res
}
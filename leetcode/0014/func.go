package leetcode

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
		max := len(strs[0])
		fmt.Println("P4")
		if len(strs) == 0 {
			return ""
		}
		if len(strs) == 1 {
			return strs[0]
		}


		for i := 0; i < max; i++ {
			if max == 0 {
				return ""
			}
			for k := 0; k < len(strs); k++ {
				if len(strs[k]) < max {
					max = len(strs[k])
				}
				if max == 0 {
					return ""
				}
				if strs[k][0: i+1] != strs[0][0: i+1] {
					return strs[0][0: i]
				}
			}
		}
		return strs[0][0: max]
}
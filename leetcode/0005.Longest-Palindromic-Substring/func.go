package leetcode

import (
	"fmt"
)

func longestPalindrome(s string) string {
		if len(s) <2 {
			return s
		}
		max := 1
		var res string
		sb := []byte(s)
		r := []byte(s)
		fmt.Println(sb)
    for i := 0; i < len(sb); i++ {
			r[i] = sb[len(sb) -i -1]
		}
		for i := 0; i < len(sb); i++ {
			for b := (i+max); b <= len(sb); b++ {
				if string(sb[i:b]) == string(r[len(r)-b:len(r)-i]) {
					fmt.Println(string(sb[i:b]), string(r[len(r)-b:len(r)-i]), i, b, len(r) )
					max = b - i+1
					res = string(sb[i:b])
				}
			}
		}
		return res
}
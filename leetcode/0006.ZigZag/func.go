package leetcode 
import (
	"fmt"
)

func convert(s string, numRows int) string {
	res := ""
	if numRows == 1 {
		return s
} 
	m := make(map[int]string)
	for i := 1; i<=numRows; i++ {
		m[i] = ""
	}
	p := 1
	k := 1
	for i := 0; i<len(s); i++ {
		fmt.Println(string(s[i]))
		m[p] = m[p] + string(s[i])
		p = p+k
		if p == 1 {
			k = 1
		} 
		if p == numRows {
			k = -1
		}
	} 
	for i := 1; i<=numRows; i++ {
		res = res + m[i]
	}
	fmt.Println(m)
	return res
}
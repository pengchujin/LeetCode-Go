package leetcode

import (
	"fmt"
	"testing"
)

type question4 struct {
	para4
	ans4
}

// para 是参数
// one 代表第一个参数
type para4 struct {
	nums1 string
	nums2 int
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one string
}

func Test_Problem4(t *testing.T) {

	qs := []question4{

		question4{
			para4{"AB", 2},
			ans4{"AB"},
		},

		question4{
			para4{"PAYPALISHIRING", 3 },
			ans4{"bbbbbbb"},
		},
		question4{
			para4{"a", 3},
			ans4{"a"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		_, p := q.ans4, q.para4
		fmt.Printf("【input】:%v       【output】:%v\n", p, convert(p.nums1, p.nums2))
	}
	fmt.Printf("\n\n\n")
}

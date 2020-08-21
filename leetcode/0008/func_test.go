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
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one int
}

func Test_Problem4(t *testing.T) {

	qs := []question4{

		question4{
			para4{"abc2abcbb"},
			ans4{2},
		},

		question4{
			para4{"abcabcbbbbbbb"},
			ans4{2},
		},
		question4{
			para4{"a"},
			ans4{2},
		},
		question4{
			para4{"ab"},
			ans4{2},
		},
		question4{
			para4{"bbbbb"},
			ans4{2},
		},
		question4{
			para4{"bb"},
			ans4{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		_, p := q.ans4, q.para4
		fmt.Printf("【input】:%v       【output】:%v\n", p, myAtoi(p.nums1))
	}
	fmt.Printf("\n\n\n")
}

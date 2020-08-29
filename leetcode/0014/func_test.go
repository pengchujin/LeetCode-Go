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
	nums1 []string
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one string
}

func Test_Problem14(t *testing.T) {

	qs := []question4{

		question4{
			para4{[]string{"asdfjikds", "a"}},
			ans4{"12"},
		},

		question4{
			para4{[]string{"123", "12345"}},
			ans4{"12"},
		},
		question4{
			para4{[]string{"123", "12345"}},
			ans4{"12"},
		},
		question4{
			para4{[]string{"123", "12345"}},
			ans4{"12"},
		},
		question4{
			para4{[]string{"123", "12345"}},
			ans4{"12"},
		},
		question4{
			para4{[]string{"123", "12345"}},
			ans4{"12"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		_, p := q.ans4, q.para4
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestCommonPrefix(p.nums1))
	}
	fmt.Printf("\n\n\n")
}

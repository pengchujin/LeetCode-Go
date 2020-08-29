package leetcode

import (
	"fmt"
)

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

func letterCombinations(digits string) []string {	
	if len(digits) < 1 {
		return res
	}
	tmp := []string{}
	for k := 0; k < len(letterMap[digits[0] - '0']); k++ {
		tmp = append(tmp, string(letterMap[digits[0] - '0'][k]))
	}

	if len(digits) == 1 {
		return tmp
	}
	
	for j := 1; j < len(digits) ; j++ {
		fmt.Println(j)
		t := []string{}
		fmt.Println(tmp)
		for i := 0; i < len(tmp); i++ {
			for k := 0; k < len(letterMap[digits[j] - '0']); k++ {
				t = append(t, string(tmp[i]) + string(letterMap[digits[j] - '0'][k]))
				fmt.Println(t, " ?")
			}	
		}
		tmp = t
	}
	return tmp
}

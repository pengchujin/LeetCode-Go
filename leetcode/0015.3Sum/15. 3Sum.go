package leetcode

import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	zheng := map[int]int{}
	fu := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			if _, ok := zheng[nums[i]]; ok {
				zheng[nums[i]] = zheng[nums[i]] + 1
			} else {
				zheng[nums[i]] = 1
			}
		} 
		if nums[i] < 0 {
			if _, ok := fu[-nums[i]]; ok {
				fu[-nums[i]] = fu[-nums[i]] + 1
			} else {
				fu[-nums[i]] = 1
			}
		} 
	}
	ze := 0
	if _, ok := zheng[0]; ok {
		ze = zheng[0]
	}
	tmp := []int{}
	tmp1 := []int{}
	if ze > 2 {
		tmp = append(tmp1, 0, 0, 0)
		res = append(res, tmp)
	}  
	for k, _ := range zheng {
		if ze > 0  {
			if _, ok := fu[k]; ok {
				tmp = append(tmp1,-k, 0, k)
				res = append(res, tmp)
			}
		} 
			for k1, _ := range fu {
				val := k - k1
				if val > 0 && k != 0 {
					if _, ok := fu[val]; ok {
						if 2*val == k && fu[val] >= 2 {
							tmp = append(tmp1,-k1, -val, k)
							res = append(res, tmp)
						}
						if 2*val != k {
							tmp = append(tmp1,-k1, -val, k)
							sort.Ints(tmp)
								res = append(res, tmp)
						}
					}
				}
				if val < 0 && k != 0{
					if _, ok := zheng[-val]; ok {
						if -val == k && zheng[k] >= 2 {
							tmp = append(tmp1,-k1, -val, k)
							res = append(res, tmp)
						}
						if -val != k {
							tmp = append(tmp1,-k1, -val, k)
							sort.Ints(tmp)
							res = append(res, tmp)
							fmt.Println(-k1, -val, k, "fuuu")
						}
					}
				}
		}
		delete(zheng, k)
	}

	fmt.Println(zheng, fu)
	return res
}

package main

import (
	"fmt"
)

func main() {
	i := thirdMax([]int{1, 1, 2, 8, 9, 7, 4})
	fmt.Println(i)
}

func thirdMax(nums []int) int {
	ch := make(chan int)
	i, j := 0, 0
	for {
		if i >= len(nums)-1 {
			break
		}
		for j = i + 1; j < len(nums) && nums[i] == nums[j]; j++ {
		}
		nums = append(nums[:i+1], nums[j:]...)
		i++
	}
	if len(nums) == 2 {
		return nums[1]
	}
	if len(nums) == 1 {
		return nums[0]
	}
	fmt.Println(nums)
	k := 0
	go func() {
		tem := nums[0]
		for {
			for k = 1; k < len(nums); k++ {
				if tem < nums[k] {
					tem = nums[k]
				}
			}

			ch <- tem
		}
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	return <-ch
}

//func thirdMax(nums []int) int {
//	//sort.Ints(nums)
//	quickSort(0, len(nums)-1, nums)
//	i, j := 0, 0
//	for {
//		if i >= len(nums)-1 {
//			break
//		}
//		for j = i + 1; j < len(nums) && nums[i] == nums[j]; j++ {
//		}
//		nums = append(nums[:i+1], nums[j:]...)
//		i++
//	}
//	if len(nums) == 2 {
//		return nums[1]
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	return nums[len(nums)-3]
//}

//type Slice []int
//
//func main() {
//	i := thirdMax([]int{1, 1, 3, 4, 2})
//	fmt.Println(i)
//}
//
//func thirdMax(nums []int) int {
//	sort.Sort(Slice(nums))
//	i, j := 0, 0
//	for {
//		if i >= len(nums)-1 {
//			break
//		}
//		for j = i + 1; j < len(nums) && nums[i] == nums[j]; j++ {
//		}
//		nums = append(nums[:i+1], nums[j:]...)
//		i++
//	}
//	if len(nums) <= 2 {
//		return nums[0]
//	}
//	return nums[2]
//
//}
//func (s Slice) Len() int {
//	return len(s)
//}
//func (s Slice) Less(i, j int) bool {
//	return s[i] > s[j] //递减
//}
//func (s Slice) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}

func quickSort(left int, right int, array []int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		quickSort(left, r, array)
	}
	if right > l {
		quickSort(l, right, array)
	}
}

//414. 第三大的数
//给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
//
//
//
//示例 1：
//
//输入：[3, 2, 1]
//输出：1
//解释：第三大的数是 1 。
//示例 2：
//
//输入：[1, 2]
//输出：2
//解释：第三大的数不存在, 所以返回最大的数 2 。
//示例 3：
//
//输入：[2, 2, 3, 1]
//输出：1
//解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
//此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。

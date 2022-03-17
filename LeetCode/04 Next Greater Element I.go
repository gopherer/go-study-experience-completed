package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	tem := []int{}
	//sort.Ints(nums2)
	// for _, v1 := range nums1{
	//     for _, v2 := range nums2{
	//         if v1 < v2 {
	//             tem = append(tem,v2)
	//             break
	//         }
	//     }
	//     tem = append(tem,-1)
	// }
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			// if i+j > len(nums2){
			//     break
			// }
			t := 1
			if nums1[i] == nums2[j] {
				for {
					if t+j >= len(nums2) {
						tem = append(tem, -1)
						break
					}
					if nums1[i] < nums2[j+t] {
						tem = append(tem, nums2[j+t])
						break
					}
					t++
				}
			}
		}
	}
	return tem
}

//0ms 例子
//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	tmp := map[int]int{}
//	for i, v := range nums2 {
//		tmp[v] = i
//	}
//
//	out := []int{}
//	for _, v := range nums1 {
//		tmpnum := -1
//		for i := tmp[v]; i < len(nums2); i++ {
//			if nums2[i] > v {
//				tmpnum = nums2[i]
//				break
//			}
//		}
//
//		out = append(out, tmpnum)
//	}
//
//	return out
//}
//496. 下一个更大元素 I
//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//
//
//示例 1:
//
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//示例 2:
//
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//
//
//提示：
//
//1 <= nums1.length <= nums2.length <= 1000
//0 <= nums1[i], nums2[i] <= 104
//nums1和nums2中所有整数 互不相同
//nums1 中的所有整数同样出现在 nums2 中

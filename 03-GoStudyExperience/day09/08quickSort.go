package main

import "fmt"

func main() {
	arr := [6]int{5, -3, 78, 45, 27, 99}
	quickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
func quickSort(left int, right int, array *[6]int) {
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

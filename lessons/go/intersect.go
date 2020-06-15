package main

import "fmt"

func main() {
	var nums1 = []int{1,2,3,4,1}
	var nums2 = []int{2,3,4,5}

	var result = intersect(nums1, nums2)
	fmt.Println(result)
}

//找出数组公共的部分
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var arr []int
	//用map存nums1中值出现次数
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		times, ok := m[v]
		if ok && times > 0 {
			arr = append(arr, v)
			m[v]--
		}
	}
	return arr
}
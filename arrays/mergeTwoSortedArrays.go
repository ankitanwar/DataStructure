package array

import (
	"fmt"
	"math"
)

//MergeTwoSortedArray : Given two sorted array we need to merge them without using extra space
func MergeTwoSortedArray(nums1, nums2 []int) {
	if len(nums2) == 0 {
		return
	}
	j := 0
	if len(nums1) > len(nums2) {
		for i := 0; i < len(nums1); i++ {
			if j < len(nums2) && nums1[i] > nums2[j] {
				for k := len(nums1) - 2; k >= i; k-- {
					nums1[k+1] = nums1[k]
				}
				println("the value of i and k is ", i, j)
				nums1[i] = nums2[j]
				j++
			}
		}
	}
	if j <= len(nums2)-1 {
		t := len(nums2) - 1
		t2 := len(nums1) - 1
		for t >= j {
			nums1[t2] = nums2[t]
			t--
			t2--
		}
	}
}

//MergeTwoSortedArray2 : Given two sorted arrays we need to merge them without using extra space
func MergeTwoSortedArray2(nums1, nums2 []int) {
	//we will use the gap strategies for these
	gap := 2
	fmt.Println("The value of gap is ", gap, len(nums1), len(nums2))
	for gap > 0 {
		start := 0
		for start+gap < len(nums1) {
			if nums1[start] > nums1[start+gap] {
				nums1[start], nums1[start+gap] = nums1[start+gap], nums1[start]
			}
			start++
		}
		var j int
		if gap > len(nums1) {
			j = gap - len(nums1) - 1
			for start < len(nums1) && j < len(nums2) {
				if nums1[start] > nums2[j] {
					nums1[start], nums2[j] = nums2[j], nums1[start]
				}
				start++
				j++
			}
		}
		if j < len(nums2) {
			j = 0
			for j+gap < len(nums2) {
				nums2[j], nums2[j+gap] = nums2[j+gap], nums2[j]
				j++
			}
		}
		gap = int(math.Ceil(float64((gap) / 2)))

	}
	fmt.Println(nums1, nums2)
}

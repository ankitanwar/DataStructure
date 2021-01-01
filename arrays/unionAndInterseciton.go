package array

import "fmt"

//UnionOfTwoArray : Given two sorted array we need to find the union of those arrays
func UnionOfTwoArray(arr, arr2 []int) {
	union := []int{}
	f := 0
	s := 0
	for f < len(arr) && s < len(arr2) {
		if arr[f] < arr2[s] {
			if f+1 < len(arr) && arr[f+1] == arr[f] {
				f++
			} else if f+1 < len(arr) && arr[f+1] != arr[f] {
				union = append(union, arr[f])
				f++
			} else {
				union = append(union, arr[f])
			}
		} else if arr2[s] < arr[f] {
			if s+1 < len(arr2) && arr2[s+1] == arr[s] {
				s++
			} else if s+1 < len(arr2) && arr2[s+1] != arr2[s] {
				union = append(union, arr2[s])
				s++
			} else {
				union = append(union, arr2[s])
			}
		} else {
			if s+1 < len(arr2) && arr2[s+1] == arr[s] {
				s++
			} else if f+1 < len(arr) && arr[f+1] == arr[f] {
				f++
			} else {
				union = append(union, arr2[s])
				f++
				s++
			}
		}
	}
	for f < len(arr) {
		if f+1 < len(arr) && arr[f] == arr[f+1] {
			f++
		} else {
			union = append(union, arr[f])
		}
	}
	for s < len(arr2) {
		if s+1 < len(arr2) && arr2[s] == arr2[s+1] {
			s++
		} else {
			union = append(union, arr2[s])
			s++
		}
	}
	fmt.Println(union)

}

//Intersection : Given two sorted array we need to find the intersection of these arrays
func Intersection(arr1, arr2 []int) {
	intersection := []int{}
	f := 0
	s := 0
	for f < len(arr1) && s < len(arr2) {
		if arr1[f] < arr2[s] {
			f++
		} else if arr1[f] > arr2[s] {
			s++
		} else {
			if f+1 < len(arr1) && arr1[f] == arr1[f+1] {
				f++
			} else if s+1 < len(arr2) && arr2[s] == arr2[s+1] {
				s++
			} else {
				intersection = append(intersection, arr1[f])
				f++
				s++
			}
		}
	}
	fmt.Println(intersection)
}

package heaps

import "container/heap"

//TopKFrequent : To find the top k frequent numbers in the given array
func TopKFrequent(arr []int, k int) []int {
	ans := []int{}
	h := &MinKeyHeap{}
	heap.Init(h)

	dict := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		current := arr[i]
		_, found := dict[current]
		if found {
			dict[current]++
		} else {
			dict[current] = 1
		}
	}

	for keys, values := range dict {
		temp := &KeyValue{
			Key:   values,
			Value: keys,
		}

		heap.Push(h, temp)
		if h.Len() > k {
			heap.Pop(h)
		}

	}

	for {
		if h.Len() == 0 {
			break
		}
		current := h.Pop().(*KeyValue).Value
		ans = append(ans, current)
	}

	return ans
}

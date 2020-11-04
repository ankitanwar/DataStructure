package heaps

import "container/heap"

type elements struct {
	value      int
	row        int
	nextcolumn int
}

type array []*elements

func (h array) Less(i, j int) bool { return h[i].value < h[j].value }
func (h array) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h array) Len() int           { return len(h) }

func (h *array) Pop() interface{} {
	current := *h
	ans := current[len(current)-1]
	*h = current[:len(current)-1]

	return ans
}

func (h *array) Push(value interface{}) {
	item := value.(*elements)
	*h = append(*h, item)
}

//SmallestRangeInKList : To determine the smallest range in the given k list
func SmallestRangeInKList(matrix [][]int) int {
	ans := 999999

	h := &array{}
	heap.Init(h)

	max := -99999

	for i := 0; i < len(matrix); i++ {
		temp := &elements{
			value:      matrix[i][0],
			row:        i,
			nextcolumn: 1,
		}
		if max < matrix[i][0] {
			max = matrix[i][0]
		}
		heap.Push(h, temp)
	}

	for {
		current := heap.Pop(h)
		ans = min(ans, max-current.(*elements).value)

		if current.(*elements).nextcolumn < len(matrix[current.(*elements).row]) {
			temp := &elements{
				value:      matrix[current.(*elements).row][current.(*elements).nextcolumn],
				row:        current.(*elements).row,
				nextcolumn: current.(*elements).nextcolumn + 1,
			}
			if temp.value > max {
				max = temp.value
			}
			heap.Push(h, temp)
		} else {
			break
		}

	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

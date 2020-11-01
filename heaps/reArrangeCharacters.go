package heaps

// import (
// 	"container/heap"
// )

// //RearrangeCharacters : to re arrange characters such that no two elements are adjacent
// func RearrangeCharacters(s string) string {
// 	dict := make(map[string]int)

// 	for i := 0; i < len(s); i++ {
// 		char := string(s[i])
// 		_, found := dict[char]
// 		if found {
// 			dict[char]++
// 		} else {
// 			dict[char] = 1
// 		}
// 	}

// 	h := &MaxKeyHeap{}
// 	heap.Init(h)

// 	for keys, values := range dict {
// 		temp := &KeyValuee{
// 			Key:   keys,
// 			Value: values,
// 		}
// 		heap.Push(h, temp)

// 	}

// 	res := ""

// 	for {
// 		if h.Len() == 1 {
// 			current := heap.Pop(h)
// 			if current.(*KeyValuee).Value > 1 {
// 				return "Not possible"
// 			}
// 			res += current.(*KeyValuee).Key
// 			return res
// 		}
// 		first := heap.Pop(h)
// 		second := heap.Pop(h)
// 		res += first.(*KeyValuee).Key
// 		res += second.(*KeyValuee).Key
// 		second.(*KeyValuee).Value--
// 		first.(*KeyValuee).Value--
// 		if first.(*KeyValuee).Value > 0 {
// 			heap.Push(h, first)
// 		}
// 		if second.(*KeyValuee).Value > 0 {
// 			heap.Push(h, second)
// 		}
// 	}

// }

package heaps

//MergeTwoMaxHeaps  : It will merge the two max heaps
func MergeTwoMaxHeaps(h1, h2 *MaxHeap) *Heaps {
	ans := Newheaps(len(*h1) + len(*h2))
	for i := 0; i < h1.Len(); i++ {
		ans.Insert((*h1)[i])
	}
	for j := 0; j < h2.Len(); j++ {
		ans.Insert((*h2)[j])
	}

	ans.BuildMaxHeap()
	ans.Print()
	return ans

}

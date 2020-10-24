package stackandqueues

import "fmt"

//Kstacks : Implement k stacks in an array
type Kstacks struct {
	arr  []int
	top  []int
	next []int
	free int
}

//Init : Constructor for the kstack
func (k *Kstacks) Init(size, n int) {
	k.arr = make([]int, size)
	k.top = make([]int, n)
	k.next = make([]int, size)
	for i := 0; i < size-1; i++ {
		k.next[i] = i + 1
	}
	k.next[size-1] = -1
	k.free = 0

}

//Insert : To insert the values in an k stack
func (k *Kstacks) Insert(val, n int) string {
	if k.free == -1 {
		return "stack overflow"
	}
	k.arr[k.free] = val
	i := k.free
	k.free = k.next[k.free]
	k.next[i] = k.top[n-1]
	k.top[n-1] = i
	return "element has been added to the database successfully"
}

//GetTop : To get the top value of given array n
func (k *Kstacks) GetTop(n int) int {
	if k.top[n] == -1 {
		return -1
	}
	return k.arr[k.top[n]]
}

//Print : To the print the values in the next
func (k *Kstacks) Print() {
	println("The value of next is ")
	for i := 0; i < len(k.next); i++ {
		fmt.Print(k.next[i])
		fmt.Print(" ")
	}
	println("")
	println("The value of top is ")
	for i := 0; i < len(k.top); i++ {
		fmt.Print(k.top[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

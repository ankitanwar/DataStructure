package backtracking

import "fmt"

var res int

//PartitionIntoKSubset : We have to partion the given string into non empty k subset
func PartitionIntoKSubset(given string, k int) {
	partions := make([][]string, k)
	partionHelper(given, k, 0, &partions)
	fmt.Println(res)
}

func partionHelper(given string, k, currentIndex int, partions *[][]string) {
	if currentIndex == len(given) {
		var check bool
		for i := 0; i < len(*partions); i++ {
			if (*partions)[i] == nil {
				check = true
				break
			}
		}
		if check != true {
			res++
			for i := 0; i < len(*partions); i++ {
				fmt.Printf("%v ", (*partions)[i])
			}
			fmt.Println("")
		}
		return
	}
	for i := 0; i < len(*partions); i++ {
		if (*partions)[i] != nil {
			(*partions)[i] = append((*partions)[i], string(given[currentIndex]))
			partionHelper(given, k, currentIndex+1, partions)
			(*partions)[i] = (*partions)[i][:len((*partions)[i])-1]
		} else {
			(*partions)[i] = append((*partions)[i], string(given[currentIndex]))
			partionHelper(given, k, currentIndex+1, partions)
			(*partions)[i] = (*partions)[i][:len((*partions)[i])-1]
			break
		}
	}
}

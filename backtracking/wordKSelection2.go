package backtracking

import "fmt"

//WordKSelection2 : We have to find out the number of ways to select k words from the strings
func WordKSelection2(given string, k int) {
	selection2Helper(given, "", 0, -1, k)
}

//lastselection ke next se choices ka pata chalega kaha character rkhna hai lastselection ke aage he characters place ho paayenge to avoid duplicates
func selection2Helper(given, ans string, selected, lastSelection, k int) {
	if selected == k {
		fmt.Println(ans)
		return
	}
	for i := lastSelection + 1; i < len(given); i++ {
		current := string(given[i])
		selection2Helper(given, ans+current, selected+1, i, k)
	}
}

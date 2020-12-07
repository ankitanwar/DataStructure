package backtracking

import "fmt"

//Combination : We have to print all the combinations in which we can select the given string into k boxes
func Combination(given string, k int) {
	helperCombination(given, "", k, 0, 0)
}

func helperCombination(given, ans string, k, currentIndex, selected int) {
	if currentIndex == len(given) {
		if selected == k {
			fmt.Println(ans)
		}
		return
	}
	helperCombination(given, ans, k, currentIndex+1, selected)
	helperCombination(given, ans+string(given[currentIndex]), k, currentIndex+1, selected+1)

}

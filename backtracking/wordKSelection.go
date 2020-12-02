package backtracking

import "fmt"

//WordKSelection : We have given a string and we need to find out the numbers of ways in which we can print the unique character of size k from given string
func WordKSelection(given string, k int) {
	getUniqueString := getUnique(given)
	helperKSelection(getUniqueString, 0, 0, k, "")
}

func helperKSelection(given string, selected, currentIndex, k int, ans string) {
	if currentIndex == len(given) {
		if selected == k {
			fmt.Println(ans)
		}
		return
	}
	curr := string(given[currentIndex])
	helperKSelection(given, selected+1, currentIndex+1, k, ans+curr)
	helperKSelection(given, selected, currentIndex+1, k, ans)

}

func getUnique(given string) string {
	dict := make(map[string]bool)
	unique := ""
	for i := 0; i < len(given); i++ {
		_, found := dict[string(given[i])]
		if !found {
			unique += string(given[i])
			dict[string(given[i])] = true

		}
	}
	return unique
}

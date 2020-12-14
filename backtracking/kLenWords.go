package backtracking

import "fmt"

//KlenWords : How many different words can we form of k size from the given string
func KlenWords(s string, k int) {
	u := getUnique(s)
	spots := make([]string, k)
	helperKLenwords(u, 0, 0, k, &spots)
}

func helperKLenwords(s string, selected, currentIndex, k int, spots *[]string) {
	if currentIndex == len(s) {
		if selected == k {
			fmt.Println(spots)
		}
		return
	}
	for i := 0; i < len(*spots); i++ {
		if (*spots)[i] == "" {
			(*spots)[i] = string(s[currentIndex])
			helperKLenwords(s, selected+1, currentIndex+1, k, spots)
			(*spots)[i] = ""
		}
	}
	helperKLenwords(s, selected, currentIndex+1, k, spots)
}

package backtracking

import "fmt"

//Klenwords2 : To calc the number of words that can be formed from the s string of len k
func Klenwords2(s string, k int) {
	used := make(map[string]bool)
	u := getUnique(s)
	helperKLenWords2(u, "", 0, k, &used)
}

func helperKLenWords2(s, asf string, selected, k int, used *map[string]bool) {
	if selected == k {
		fmt.Println(asf)
		return
	}
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		_, found := (*used)[current]
		if !found {
			(*used)[current] = true
			helperKLenWords2(s, asf+current, selected+1, k, used)
			delete(*used, current)
		}
	}
}

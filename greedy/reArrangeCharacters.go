package greedy

import (
	"fmt"
	"sort"
)

//KeyValue : Key value pair for a dictonary
type KeyValue struct {
	key   string
	value int
}

//RearrangeCharacter : we are given a string and we have to rearrange the strings such that no alternate in the string is alternate
func RearrangeCharacter(s string) {
	fmt.Println(len(s))
	dict := make(map[string]int)
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		_, found := dict[current]
		if !found {
			dict[current] = 1
		} else {
			dict[current]++
		}
	}
	if len(dict) < 2 {
		fmt.Println("Not Possible")
	} else {
		keys := []KeyValue{}
		for k, v := range dict {
			keys = append(keys, KeyValue{k, v})
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i].value > keys[j].value })
		fmt.Println(keys)
		first := keys[0].key
		second := keys[1].key
		i := 2
		var ans string
		for {
			if dict[first] == 0 {
				if i < len(keys) {
					first = keys[i].key
					fmt.Println(first)
					i++
				} else {
					break
				}
			}
			if dict[second] == 0 {
				if i < len(keys) {
					second = keys[i].key
					fmt.Println(second)
					i++
				} else {
					break
				}
			}
			ans += first
			ans += second
			dict[first]--
			dict[second]--
		}
		if dict[first] != 0 && string(ans[len(ans)-1]) != first {
			ans += first
			dict[first]--
		} else if dict[second] != 0 && string(ans[len(ans)-1]) != second {
			ans += second
			dict[second]--
		}
		if dict[first] != 0 && string(ans[0]) != first {
			ans = first + ans
			dict[first]--
		} else if dict[second] != 0 && string(ans[0]) != second {
			ans = second + ans
			dict[second]--
		}
		fmt.Println(dict[first])
		fmt.Println(dict[second], second)
		if len(ans) == len(s) {
			fmt.Println(ans)
		} else {
			fmt.Println()
		}
	}
}

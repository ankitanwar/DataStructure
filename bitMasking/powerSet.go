package bit

import "fmt"

//PowerSet : Given a string we need to generate all of its power set
func PowerSet(s string) {
	r := 1 << len(s)
	for i := 0; i < r; i++ {
		for j := 0; j < len(s); j++ {
			t := 1 << j
			if i&t != 0 {
				fmt.Print(string(s[j]))
			}
		}
		println()
	}
}

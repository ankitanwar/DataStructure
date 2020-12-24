package graph

import (
	"fmt"
	"math"
)

//AlienDictonary : Given a list of words and we need to determine the order of the words
func AlienDictonary(words []string) {
	dict := make(map[string][]string)
	for i := 0; i < len(words); i++ {
		current := string(words[i])
		for j := 0; j < len(current); j++ {
			_, found := dict[string(current[j])]
			if !found {
				dict[string(current[j])] = []string{}
			}

		}
	}
	inEdges := make(map[string]int)
	for key := range dict {
		inEdges[key] = 0
	}
	for i := 1; i < len(words); i++ {
		first := words[i-1]
		second := words[i]
		l := int(math.Min(float64(len(first)), float64(len(second))))
		for j := 0; j < l; j++ {
			if first[j] != second[j] {
				dict[string(first[j])] = append(dict[string(first[j])], string(second[j]))
				inEdges[string(second[j])]++
				break
			}
		}
	}
	q := []string{}
	for key := range inEdges {
		if inEdges[key] == 0 {
			q = append(q, key)
		}
	}
	ans := ""
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		ans += removed
		current := dict[removed]
		for i := 0; i < len(current); i++ {
			inEdges[string(current[i])]--
			if inEdges[string(current[i])] == 0 {
				q = append(q, string(current[i]))
			}
		}

	}
	fmt.Println(ans)
}

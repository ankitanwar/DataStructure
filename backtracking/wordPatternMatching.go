package backtracking

import "fmt"

//pattern ->    p   e      p
//words ->   graph tree graph

//WordPatternMatching : we need to map every pattern with the given string
func WordPatternMatching(pattern, words, originalPattern string, dict *map[string]string) {
	if len(pattern) == 0 {
		if len(words) == 0 {
			alreadyPrinted := make(map[string]bool) // to avoid duplicate printing
			for i := 0; i < len(originalPattern); i++ {
				current := string(originalPattern[i])
				_, found := alreadyPrinted[current]
				if found == false {
					fmt.Println(current, (*dict)[current])
					alreadyPrinted[current] = true
				}
			}
		}
		return
	}

	currentPattern := string(pattern[0])
	pattern = pattern[1:]
	_, found := (*dict)[currentPattern]
	if found {
		previousMapped := (*dict)[currentPattern]
		if len(words) >= len(previousMapped) {
			check := words[:len(previousMapped)]
			right := words[len(previousMapped):]
			if check == previousMapped {
				WordPatternMatching(pattern, right, originalPattern, dict)
			}
		}

	} else {
		for j := 0; j < len(words); j++ {
			left := words[0 : j+1]
			right := words[j+1:]
			(*dict)[currentPattern] = left
			WordPatternMatching(pattern, right, originalPattern, dict)
			delete(*(dict), currentPattern)

		}
	}
}

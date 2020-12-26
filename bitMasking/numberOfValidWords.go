package bit

import (
	"fmt"
)

//NumberOfValidWords : Given an array of words and puzzle we need to find out the number of valid words corresponding to each puzzle
func NumberOfValidWords(words, puzzle []string) {
	arr := []int{}
	for i := 0; i < len(puzzle); i++ {
		arr = append(arr, 0)
	}
	dict := make(map[string][]int)
	t := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(t); i++ {
		dict[string(t[i])] = []int{}
	}
	for i := 0; i < len(words); i++ {
		current := string(words[i])
		b := 0
		unique := make(map[string]bool)
		for j := 0; j < len(current); j++ {
			_, found := unique[string(current[j])]
			if !found {
				b = b | (1 << (rune(current[j]) - 'a'))
				unique[string(current[j])] = true
			}
		}
		for key := range unique {
			dict[key] = append(dict[key], b)
		}
	}
	for i := 0; i < len(puzzle); i++ {
		cb := 0
		currentPuzzle := string(puzzle[i])
		unique := make(map[string]bool)
		for j := 0; j < len(currentPuzzle); j++ {
			_, found := unique[string(currentPuzzle[j])]
			if !found {
				unique[string(currentPuzzle[j])] = true
				cb = cb | 1<<(rune(currentPuzzle[j])-'a')
			}
		}
		possibles := dict[string(currentPuzzle[0])]
		if len(possibles) != 0 {
			for k := 0; k < len(possibles); k++ {
				if cb&possibles[k] == possibles[k] {
					arr[i] = arr[i] + 1
				}
			}
		}
	}
	fmt.Println(arr)
}

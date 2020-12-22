package graph

type ladderLength struct {
	levl int
	word string
}

//LadderLength : Given a start word and the end word and we have to count the minimum numbers requires to covert the start word into the end word
func LadderLength(beginWord string, endWord string, wordList []string) int {
	q := []ladderLength{}
	dict := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		current := string(wordList[i])
		dict[current] = true
	}
	_, found := dict[endWord]
	if !found {
		return 0
	}
	_, foundd := dict[beginWord]
	if foundd {
		dict[beginWord] = false
	}
	t := ladderLength{
		levl: 1,
		word: beginWord,
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		if removed.word == endWord {
			return removed.levl
		}
		current := removed.word
		for i := 0; i < len(current); i++ {
			aToz := "abcdefghijklmnopqrstuvwxsyz"
			for j := 0; j < len(aToz); j++ {
				start := current[:i]
				t := string(aToz[j])
				end := current[i+1:]
				newWord := start + t + end
				_, found := dict[newWord]
				if found && dict[newWord] == true {
					t := ladderLength{
						levl: removed.levl + 1,
						word: newWord,
					}
					dict[newWord] = false
					q = append(q, t)
				}

			}
		}

	}
	return 0

}

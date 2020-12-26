package bit

import "strconv"

var ans []int

//MinimumNumberOfDevelopers : To calc the minimum number of developers required to finish the job
func MinimumNumberOfDevelopers(skills []string, devs map[int][]string) {
	ans = make([]int, 1e7)
	t := ""
	skill := make(map[string]int)
	for i := 0; i < len(skills); i++ {
		skill[skills[i]] = 0
		t += "1"
	}
	newDevs := make(map[int]int)
	for key := range devs {
		jobs := devs[key]
		t2 := 0
		for i := 0; i < len(jobs); i++ {
			t3 := skill[string(jobs[i])]
			t2 = t2 | 1<<t3
		}
		newDevs[key] = t2

	}
	wanted, _ := strconv.Atoi(t)
	minDeveloper(0, len(newDevs), 0, wanted, ans, newDevs)

}

func minDeveloper(currentDev, n, currentAns, wanted int, currentIncluded []int, devs map[int]int) {
	if currentDev == n {
		if currentAns&wanted == wanted && len(currentIncluded) < len(ans) {
			currentIncluded = ans
		}
		return
	}
	minDeveloper(currentDev+1, n, currentAns, wanted, currentIncluded, devs)
	currentIncluded = append(currentIncluded, currentDev)
	minDeveloper(currentDev+1, n, currentAns|devs[currentDev], wanted, currentIncluded, devs)

}

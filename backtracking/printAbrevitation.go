package backtracking

import "strconv"

// Input : ABC
// Output :
// ABC
// AB1
// A1C
// A2
// 1BC
// 1B1
// 2C
// 3

//PrintAbrevation : To print all the abrevation in of the strings
func PrintAbrevation(given, ans string, currentPos, count int) {
	if currentPos == len(given) {
		if count == 0 {
			println(ans)
		} else {
			println(ans + strconv.Itoa(count))
		}
		return
	}

	//if we want to print the character
	if count > 0 {
		PrintAbrevation(given, ans+strconv.Itoa(count)+string(given[currentPos]), currentPos+1, 0)
	} else if count <= 0 {
		PrintAbrevation(given, ans+string(given[currentPos]), currentPos+1, 0)
	}

	//if we dont want to print the character
	PrintAbrevation(given, ans, currentPos+1, count+1)

}

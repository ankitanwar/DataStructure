package bit

import "fmt"

//ConvertAToB : GIven a number we need to convert num1 to num2
func ConvertAToB(num1, num2 int) {
	differentBit := num1 ^ num2
	var count int
	for differentBit > 0 {
		rmsb := differentBit & -differentBit
		differentBit -= rmsb
		count++
	}
	fmt.Println(count)
}

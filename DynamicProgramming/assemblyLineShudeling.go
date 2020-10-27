package dynamicprogramming

import "fmt"

// AssemblyLineShedulign : To cal the minimum time required to build an vehicle
func AssemblyLineShedulign(inTime, transferTime [][]int, startTime, endTime []int) int {
	table1 := make([]int, len(inTime[0]))
	table2 := make([]int, len(inTime[0]))

	table1[0] = inTime[0][0] + startTime[0]
	table2[0] = inTime[1][0] + startTime[1]

	for i := 1; i < len(table1); i++ {
		table1[i] = minimum(table1[i-1]+inTime[0][i], table2[i-1]+transferTime[1][i]+inTime[0][i])
		table2[i] = minimum(table2[i-1]+inTime[1][i], table1[i-1]+transferTime[0][i]+inTime[1][i])

	}
	for i := 0; i < len(table1); i++ {
		fmt.Printf("The values are %v %v ", table1[i], table2[i])
	}
	fmt.Println("")
	return minimum(table1[len(table1)-1]+endTime[0], table2[len(table2)-1]+endTime[1])

}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

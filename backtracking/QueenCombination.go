package backtracking

import "fmt"

//we have to place n identical (q,q,q)queens into board of nxn and print all the possible ways in which we can do so

//QueenCombination : To print all the possbile soln
func QueenCombination(n int) {
	helperQueen(0, 0, n, 0, "")
}

func helperQueen(row, column, n, queenSelected int, asf string) {
	if row == n {
		if queenSelected == n {
			fmt.Println(asf)
			println()
			println("Next Board")
		}
		return
	}
	var currentRow int = 0
	var currentColumn int = 0
	ifAceepted := ""
	ifNot := ""
	if column == n-1 {
		currentRow = row + 1
		currentColumn = 0
		ifAceepted += asf + "q\n"
		ifNot += asf + "-\n"
	} else {
		currentRow = row
		currentColumn = column + 1
		ifAceepted = asf + "q"
		ifNot = asf + "-"
	}
	helperQueen(currentRow, currentColumn, n, queenSelected+1, ifAceepted)
	helperQueen(currentRow, currentColumn, n, queenSelected, ifNot)
}

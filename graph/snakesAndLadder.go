package graph

import "fmt"

type game struct {
	step  int
	value int
}

//SnakesAndLadder : Given a board of nxn if board[i][i] doesn't contain -1 then there can be either ladder or snake and we will move to that exact same value we need to determine the minimum value required to reach nxn
func SnakesAndLadder(board [][]int) {
	n := len(board)
	visited := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		visited = append(visited, t)
	}
	snakeAndLadderHelper(n, &visited, board)
}

func snakeAndLadderHelper(n int, visited *[][]bool, board [][]int) {
	q := []game{}
	t := game{
		step:  0,
		value: 1,
	}
	q = append(q, t)
	for len(q) > 0 {
		fmt.Println(q)
		removed := q[0]
		if removed.value == n*n {
			fmt.Println(removed.step)
			break
		}
		q = q[1:]
		for i := 1; i <= 6; i++ {
			if removed.value+i > n*n {
				break
			}
			cordinates := findCoordinates(removed.value+i, n)

			row := cordinates[0]
			col := cordinates[1]
			if (*visited)[row][col] == true {
				continue
			} else {
				(*visited)[row][col] = true
				if board[row][col] != -1 {
					t := game{
						value: board[row][col],
						step:  removed.step + 1,
					}
					q = append(q, t)
				} else {
					t := game{
						value: removed.value + i,
						step:  removed.step + 1,
					}
					q = append(q, t)
				}
			}

		}
	}

}

func findCoordinates(value, n int) []int {
	fmt.Println("The receives value is ", value)
	r := n - ((value - 1) / n) - 1
	c := (value - 1) % n
	var row int
	var col int
	if r%2 != 0 {
		row = r
		col = n - 1 - c
	} else {
		row = r
		col = c
	}
	pos := []int{row, col}
	println("The value of row and col is ", row, col)
	return pos

}

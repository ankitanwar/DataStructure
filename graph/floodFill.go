package graph

import "fmt"

type floodfill struct {
	x int
	y int
}

//FloodFill : Given a matrix and starting index and we have to change all the adjacent of starting index having value same as starting index and their adjacent having same value by the given value
func FloodFill(matrix [][]int, startingIndex []int, givenValue int) {
	q := []floodfill{}
	t := floodfill{
		x: startingIndex[0],
		y: startingIndex[1],
	}
	startingValue := matrix[startingIndex[0]][startingIndex[1]]
	visited := [][]bool{}
	for i := 0; i < len(matrix); i++ {
		t := make([]bool, len(matrix[0]))
		visited = append(visited, t)
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		matrix[removed.x][removed.y] = givenValue
		dx := []int{0, 0, -1, 1}
		dy := []int{-1, +1, 0, 0}
		for i := 0; i < len(dx); i++ {
			put := canPut(matrix, startingValue, removed.x+dx[i], removed.y+dy[i])
			if put == true {
				if visited[removed.x+dx[i]][removed.y+dy[i]] == false {
					visited[removed.x+dx[i]][removed.y+dy[i]] = true
					t := floodfill{
						x: removed.x + dx[i],
						y: removed.y + dy[i],
					}
					q = append(q, t)
				}
			}
		}

	}
	fmt.Println(matrix)

}

func canPut(matrix [][]int, startingValue, row, col int) bool {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != startingValue {
		return false
	}
	return true
}

package graph

import "fmt"

type minimumStep struct {
	x    int
	y    int
	step int
}

func canMove(x, y, n int, visited [][]bool) bool {
	if x < 0 || x >= n || y < 0 || y >= n || visited[x][y] == true {
		return false
	}
	return true
}

//MinimumStepByKnight : Given a knight we need to calc the minimum number of steps required to reach the destionation
func MinimumStepByKnight(n int, currentPosition []int, targetPosition []int) {
	q := []minimumStep{}
	t := minimumStep{
		x:    currentPosition[0],
		y:    currentPosition[1],
		step: 0,
	}
	visited := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		visited = append(visited, t)
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		if removed.x == targetPosition[0] && removed.y == targetPosition[1] {
			fmt.Println(removed.step)
			return
		}
		if visited[removed.x][removed.y] == true {
			continue
		} else {
			visited[removed.x][removed.y] = true
			dx := []int{-2, -1, 1, 2, -2, -1, 1, 2}
			dy := []int{-1, -2, -2, -1, 1, 2, 2, 1}
			for i := 0; i < len(dx); i++ {
				move := canMove(dx[i]+removed.x, dy[i]+removed.y, n, visited)
				if move {
					t := minimumStep{
						x:    removed.x + dx[i],
						y:    removed.y + dy[i],
						step: removed.step + 1,
					}
					q = append(q, t)
				}
			}
		}
	}
}

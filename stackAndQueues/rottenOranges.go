package stackandqueues

//RottenOranges : Minimum time reuire to rotte all oranges
func RottenOranges(matrix [][]int) int {
	queue := [][]int{}
	time := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 2 {
				temp := []int{i, j}
				queue = append(queue, temp)
			}
		}
	}
	temp := []int{-1, -1} // To indicate that next frame is started
	queue = append(queue, temp)

	for {

		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]

		if current[0] == -1 && current[1] == -1 {
			if len(queue) > 0 {
				temp := []int{-1, -1}
				queue = append(queue, temp)
				time++
			}
			continue
		}

		//checkForBottom
		if current[0]+1 < len(matrix) {
			if matrix[current[0]+1][current[1]] == 1 {
				matrix[current[0]+1][current[1]] = 2
				temp := []int{current[0] + 1, current[1]}
				queue = append(queue, temp)
			}
		}

		//check For upper
		if current[0]-1 >= 0 {
			if matrix[current[0]-1][current[1]] == 1 {
				matrix[current[0]-1][current[1]] = 2
				temp := []int{current[0] - 1, current[1]}
				queue = append(queue, temp)
			}
		}

		//check for right
		if current[1]+1 < len(matrix[0]) {
			if matrix[current[0]][current[1]+1] == 1 {
				matrix[current[0]][current[1]+1] = 2
				temp := []int{current[0], current[1] + 1}
				queue = append(queue, temp)
			}
		}

		//check for left

		if current[1]-1 >= 0 {
			if matrix[current[0]][current[1]-1] == 1 {
				matrix[current[0]][current[1]-1] = 2
				temp := []int{current[0], current[1] - 1}
				queue = append(queue, temp)
			}
		}
	}

	flag := false

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				flag = true
				break
			}
		}
	}
	if flag == false {
		return time
	}
	return -1

}

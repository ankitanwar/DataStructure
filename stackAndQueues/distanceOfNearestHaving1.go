package stackandqueues

import "fmt"

//Nearest1InMatrix : nearest 1 in the matrix
func Nearest1InMatrix(matrix [][]int) [][]int {
	queue := [][]int{}
	vistited := [][]bool{}

	//setting up thr boolean matrix
	for i := 0; i < len(matrix); i++ {
		temp := []bool{}
		for j := 0; j < len(matrix[0]); j++ {
			temp = append(temp, false)
		}
		vistited = append(vistited, temp)
	}

	//getting the index of all the 1 in the given matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = 0
				vistited[i][j] = true
				temp := []int{i, j}
				queue = append(queue, temp)
			}
		}
	}

	for {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]

		//Checking Bottom
		if current[0]+1 < len(matrix) && vistited[current[0]+1][current[1]] == false {
			matrix[current[0]+1][current[1]] = 1 + matrix[current[0]][current[1]]
			temp := []int{current[0] + 1, current[1]}
			queue = append(queue, temp)

			vistited[current[0]+1][current[1]] = true
		}

		//Checking Top

		if current[0]-1 >= 0 && vistited[current[0]-1][current[1]] == false {
			matrix[current[0]-1][current[1]] = 1 + matrix[current[0]][current[1]]
			temp := []int{current[0] - 1, current[1]}
			queue = append(queue, temp)
			vistited[current[0]-1][current[1]] = true
		}

		//checking right
		if current[1]+1 < len(matrix[0]) && vistited[current[0]][current[1]+1] == false {
			matrix[current[0]][current[1]+1] = 1 + matrix[current[0]][current[1]]
			temp := []int{current[0], current[1] + 1}
			queue = append(queue, temp)

			vistited[current[0]][current[1]+1] = true
		}

		//Checking Left
		if current[1]-1 >= 0 && vistited[current[0]][current[1]-1] == false {
			matrix[current[0]][current[1]-1] = 1 + matrix[current[0]][current[1]]
			temp := []int{current[0], current[1] - 1}
			queue = append(queue, temp)

			vistited[current[0]][current[1]-1] = true
		}

	}

	println("The value of visited is")
	for i := 0; i < len(vistited); i++ {
		for j := 0; j < len(vistited[0]); j++ {
			fmt.Printf("%v ", vistited[i][j])
		}

		fmt.Println("")
	}

	return matrix

}

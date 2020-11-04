package dynamicprogramming

//EggDroping : to find the minimum number of attempts in which we can find the threshold floor
//Threshold floor is a floor above which egg is gonna break
func EggDroping(egg, floor int) int {
	if floor == 0 || floor == 1 {
		return floor
	}
	if egg == 0 {
		return 0
	}
	if egg == 1 {
		return floor
	}
	_min := 99999
	for k := 1; k <= floor; k++ {
		//in worst case we have to find do max number of attempts thefore finding the max value
		count := 1 + maximum(EggDroping(egg-1, k-1), EggDroping(egg, floor-k))

		//counting the min number of attempts require
		if count < _min {
			_min = count
		}
	}

	return _min
}

//EggDroping2 : to find out the minimum number of ways in which can find the threshold floor
func EggDroping2(egg, floor int) int {
	matrix := [100][100]int{}
	return eggSolve(&matrix, egg, floor)

}

func eggSolve(mat *[100][100]int, egg, floor int) int {

	if egg == 0 {
		mat[floor][egg] = 0
		return mat[floor][egg]
	}
	if floor == 0 || floor == 1 {
		mat[floor][egg] = floor
		return mat[floor][egg]
	}
	if egg == 1 {
		mat[floor][egg] = floor
		return mat[floor][egg]
	}
	if mat[floor][egg] != 0 {
		return mat[floor][egg]
	}
	_min := 99999
	for k := 1; k <= floor; k++ {
		//in worst case we have to find do max number of attempts thefore finding the max value
		count := 1 + maximum(eggSolve(mat, egg-1, k-1), eggSolve(mat, egg, floor-k))

		//counting the min number of attempts require
		if count < _min {
			_min = count
		}
	}
	mat[floor][egg] = _min

	return _min
}

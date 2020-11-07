package backtracking

//JosephProblem : Every kth person will be killed
func JosephProblem(NumberOfPerson, k int) int {
	if NumberOfPerson == 1 {
		return 0
	}
	x := JosephProblem(NumberOfPerson-1, k)
	y := (x + k) % NumberOfPerson
	return y

}

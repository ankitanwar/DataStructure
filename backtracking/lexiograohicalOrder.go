package backtracking

//LexoGraphical : It will print the lexopgraphical Order
func LexoGraphical(number int) {
	for i := 1; i <= 9; i++ {
		printLexo(i, number)
	}
}
func printLexo(n, number int) {
	if n > number {
		return
	}
	println(n)
	for i := 0; i < 10; i++ {
		printLexo(n*10+i, number)
	}
}

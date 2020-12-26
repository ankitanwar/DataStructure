package bit

import (
	"fmt"
	"math"
)

//Josephus : Josephus problem using bit manuplation only for alternate killing
func Josephus(n int) {
	p := math.Log2(float64(n))
	fmt.Println("The value of p is ", p)
	l := n - int(p)
	fmt.Println("The value of l is ", l)
	fmt.Println(2*l + 1)

}

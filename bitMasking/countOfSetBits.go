package bit

import (
	"math"
)

//CountOFSetBit : GIven a number we need to calc the the set bit in a given integer
func CountOFSetBit(n int) int {
	if n == 0 {
		return 0
	}
	var ans int
	p := int(math.Log2(float64(n)))
	up := p * int(math.Pow(2, float64(p)-1))
	down := n - int(math.Pow(2, float64(p))) + 1
	rest := n - int(math.Pow(2, float64(p)))
	ans = up + down + CountOFSetBit(rest)
	return ans
}

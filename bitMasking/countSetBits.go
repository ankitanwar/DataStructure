package bit

import "fmt"

//BinaryAndReverseBit : Given a number find its reverse
func BinaryAndReverseBit(n int) {
	var found1 = false
	var ans int
	var s string
	var j int = 0
	for i := 31; i >= 0; i-- {
		msk := int(1 << i)
		t := msk & n
		if found1 == true {
			if t != 0 {
				ans |= (1 << j)
				s += "1"
			} else {
				s += "0"
			}
			j++
		} else {
			if t != 0 {
				found1 = true
				ans = ans | (1 << j)
				s += "1"
				j++
			}
		}
	}
	fmt.Println(s)
	fmt.Println(ans)
}

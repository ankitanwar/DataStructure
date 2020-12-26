package bit

import "fmt"

//SwapAllOddsAndEvenBits : Given a Number swap all its odds and even bits
func SwapAllOddsAndEvenBits(n int64) {
	var evenMask int64
	evenMask = 0xAAAAAAAA
	var oddMask int64
	oddMask = 0x55555555

	evenBits := n & evenMask
	oddBits := n & oddMask

	evenBits >>= 1
	oddBits <<= 1

	fmt.Println(evenBits | oddBits)
}

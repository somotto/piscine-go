package main

import (
	"fmt"
)

func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = result<<1 | (oct & 1)
		oct >>= 1
	}
	return result
}

func main() {
	oct := byte(38) // 0010 0110
	fmt.Printf("%08b\n", oct)
	fmt.Printf("%08b\n", ReverseBits(oct))
}

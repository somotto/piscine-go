package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NotDecimal("0.00000000000000000000000000000000000000001"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Println(NotDecimal(""))
	fmt.Println(NotDecimal("-0000000000000000000000000019.525856"))
	fmt.Println(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	newdec := ""
	for _, v := range dec {
		if v != '.' {
			if v == '-' || (v >= '0' && v <= '9') {
				newdec += string(v)
			} else {
				return dec
			}
		}
	}
	n, _ := strconv.Atoi(newdec)
	return strconv.Itoa(n)
}

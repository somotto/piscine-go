package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(NotDecimal("0.00000000000000000000000000000000000000001"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-0000000000000000000000000019.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n\n"
	}
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
	return strconv.Itoa(n) + "\n"
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	count := 1
	current := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == current {
			count++
		} else {
			res += strconv.Itoa(count) + string(current)
			current = s[i]
			count = 1
		}
	}

	res += strconv.Itoa(count) + string(current)
	return res
}

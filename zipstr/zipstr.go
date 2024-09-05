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
	Newstr := ""
	count := 1

	for i:= 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			Newstr += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	Newstr += strconv.Itoa(count) + string(s[len(s)-1])
	return Newstr
}


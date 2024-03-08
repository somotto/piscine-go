package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	sort.Strings(args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
	fmt.Printf("%#v\n", Slice(a, 2, 1, 9, 8, 7))

}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs)== 0{
		return nil
	}
	start := nbrs[0]
	end := len(a)

	if start < 0{
		start = len(a)+ start
	}

	if len(nbrs) >1{
		end = nbrs[len(nbrs)-1]

		if end < 0{
			end = len(a)+ end
		}
		if end > len(a){
			end = len(a)
		}

		if start < 0 || start > end || start >= len(a){
			return nil
		}
	}
	return a[start : end]
}

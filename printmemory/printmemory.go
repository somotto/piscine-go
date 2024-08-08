package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexDigits := "0123456789abcdef"

	// Print the hexadecimal representation in lines of 4 bytes each
	for i := 0; i < len(arr); i += 4 {
		for j := i; j < i+4 && j < len(arr); j++ {
			z01.PrintRune(rune(hexDigits[arr[j]>>4]))
			z01.PrintRune(rune(hexDigits[arr[j]&0x0F]))
			if j != i+3 && j != len(arr)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

	 // Print the ASCII representation of all 10 bytes on a single line
	 for i := 0; i < len(arr); i++ {
        if arr[i] >= 32 && arr[i] <= 126 {
            z01.PrintRune(rune(arr[i]))
        } else {
            z01.PrintRune('.')
        }
    }
    z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

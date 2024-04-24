package main

import (
	"fmt"
)

func main() {
	offset := 3
	bit := uint8(0x44)

	fmt.Printf("%08b\n", bit) // 1111011

	target := uint8((1 << (8 - offset)))
	bit |= target

	fmt.Printf("%08b\n", bit) // 1111011

	size := 40
	fmt.Printf("%d\n", size>>3)

	fmt.Printf("%08b\n", 0xAA)
}

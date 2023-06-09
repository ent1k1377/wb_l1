package main

import "fmt"

func main() {
	var num int64 = 42
	var bitIndex uint = 5
	var bitValue int = 0

	setBit(&num, bitIndex, bitValue)
	fmt.Printf("Число после изменения: %d\n", num)
}

func setBit(n *int64, i uint, bitValue int) {
	mask := int64(1) << i

	if bitValue == 1 {
		*n |= mask
	} else {
		*n &= ^mask
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("expected one argument")
		os.Exit(1)
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("the number of bits set in %d is %d\n", n, PopCount(n))
	fmt.Printf("%d = %b\n", n, n)
}

func PopCount(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

package main

import (
	"fmt"
)

func iterate() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}

func main() {
	iterator := iterate()

	for num := 1; num <= 10; num++ {
		fmt.Println(iterator())
	}
}

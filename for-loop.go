package main

import "fmt"

func seperate_var() {
	num := 1
	for num <= 10 {
		fmt.Println(num)
		num++
	}
}

func one_line() {

	for num := 1; num <= 10; num++ {
		fmt.Println(num)
	}

}

func main() {

	fmt.Println("Using sperate var declaration")
	seperate_var()

	fmt.Println("Using one line declaration")
	one_line()
}

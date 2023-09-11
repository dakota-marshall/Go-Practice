package main

import (
	"fmt"
	"strconv"
)

func main() {
	test_map := make(map[string]int)

	test_map["v1"] = 1
	test_map["v2"] = 2
	test_map["v3"] = 3

	for key, value := range test_map {
		fmt.Printf("Key: %s - Value: %s\n", key, strconv.Itoa(value))
	}
}

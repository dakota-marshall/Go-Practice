package main

import (
	"fmt"
)

type Dog struct {
	name  string
	age   int
	breed string
}

func (dog Dog) get_age() int {
	return dog.age
}

type Cat struct {
	name  string
	age   int
	color string
}

func (cat Cat) get_age() int {
	return cat.age
}

type Animal interface {
	get_age() int
}

func average_age(animals ...Animal) float64 {

	var total_age int
	var animal_count int

	for _, animal := range animals {
		total_age += animal.get_age()
		animal_count++
	}

	var average float64 = float64(total_age) / float64(animal_count)
	return average

}

func main() {

	brutus := Dog{
		name:  "Brutus",
		age:   10,
		breed: "Mutt",
	}

	mercury := Cat{
		name:  "Mercury",
		age:   7,
		color: "Black",
	}

	fmt.Println(average_age(brutus, mercury))

}

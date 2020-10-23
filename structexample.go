package main

import (
	"fmt"
)

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{
		Name:       "maruthi",
		Model:      "2020",
		Color:      "white",
		WeightInKg: 1900,
	}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	c.Color = "Black"
	fmt.Println("Car: ", c)
}

package main

import "fmt"

func main() {
	// Untyped Constant
	const myFavLanguage = "Go programming language"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "India", 91

	const (
		employeeId string  = "indianrupee"
		salary     float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr)
}

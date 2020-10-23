package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func main() {
	pEmp := new(Employee)

	pEmp.Id = 1000
	pEmp.Name = "Sachin"

	fmt.Println(pEmp)
}

package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alexey"
	alex.lastName = "Potato"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}

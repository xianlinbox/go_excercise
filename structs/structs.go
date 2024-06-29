package main

import "fmt"

type person struct {
	firstName string;
	lastName string;
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	me := person{firstName: "John", lastName: "Doe"}
	me.updateName("Jane")
	fmt.Println(me)

	s :="aa"
	fmt.Println(*&s)

	name := "bill"
 
	namePointer := &name
	
	fmt.Println(&namePointer)
	printPointer(namePointer) 
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
 }
 
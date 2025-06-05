package main

import "fmt"

type HasName interface {
	GetName() string // contract method
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {// implementation method of HasName Interface
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
} 


func Interface() {
	person := Person{
		Name: "Mustofa",
	}
	animal := Animal {
		Name: "Kucing",
	}

	SayHello(person)
	SayHello(animal)
}

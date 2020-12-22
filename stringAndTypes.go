package main

import "fmt"


func stringThingsAndTypes() {

	type aType = uint8
	type bType uint8

	s := "This is a string"
	a := aType(s[0])
	b := bType(s[0])

	fmt.Printf("type of s[0] => %T\n", s[0])
	fmt.Printf("value of s[0] => %q\n", s[0])

	fmt.Printf("type of a => %T\n", a)
	fmt.Printf("value of a => %q\n", a)

	fmt.Printf("type of b => %T\n", b)
	fmt.Printf("value of b => %q\n", b)
}

func enumThings() {
	type Vehicle int

	const (
		Bike Vehicle = iota
		Scooter
		Car
		Bus
		Train
	)

	v := Vehicle(5)
	fmt.Printf("value v => %v\n", v)
	fmt.Printf("type v => %T\n", v)
}
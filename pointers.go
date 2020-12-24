package main

import "fmt"

type Thing struct {
	Name string
}

type Bag []Thing

func (b Bag) Find(name string) *Thing {
	for i, t := range b {
		if t.Name == name {
			return &b[i]
		}
	}

	return nil
}

func pointerThing() {

	bag := Bag([]Thing{
		{"Shoe"},
		{"Sock"},
		{"Guitar"},
		{"Fan"},
	})

	fmt.Println(bag)

	shoe := bag.Find("Shoe")
	shoe.Name = "Fancy Shoe"

	fmt.Println(bag)
}

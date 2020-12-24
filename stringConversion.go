package main

import (
	"fmt"
	"strconv"
)

func stringConversions() {
	num := 42

	s := strconv.Itoa(num)

	fmt.Printf("num => %#v\n", num)
	fmt.Printf("s => %#v\n", s)
}

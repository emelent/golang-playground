package main

import (
	"errors"
	"fmt"
)

func sentinelErrors() {
	errA := errors.New("mush mush")
	errB := errors.New("mush mush")

	fmt.Printf("errA == errB is %v\n", errA == errB)
	fmt.Printf("errA.string() == errB.string() is %v\n", errA.Error() == errB.Error())
}

const (
	MyErrorA = MyError("ErrorA")
)

type MyError string

func (e MyError) Error() string { return string(e) }

func sentinelErrors2() {

	var errA MyError = MyErrorA
	errB := MyErrorA
	errC := MyError("ErrorA")

	fmt.Printf("errA == errB is %v\n", errA == errB)
	fmt.Printf("errA == errC is %v\n", errA == errC)
	fmt.Printf("errA.string() == errB.string() is %v\n", errA.Error() == errB.Error())
	fmt.Printf("errA.string() == errC.string() is %v\n", errA.Error() == errC.Error())

	switch errA {
	case MyErrorA:
		fmt.Println("errA is MyErrorA")
	default:
		fmt.Println("errA is not MyErrorA")
	}
}


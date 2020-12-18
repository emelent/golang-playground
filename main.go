package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func printLen(i []int) {
	fmt.Printf("len => %d\n", len(i))
}

func printCap(i []int) {
	fmt.Printf("cap => %d\n", cap(i))
}

func sliceFun() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceA := arr[:]
	sliceB := arr[1:3]

	fmt.Printf("arr => %v\n", arr)
	fmt.Printf("sliceA => %v\n", sliceA)
	printLen(sliceA)
	printCap(sliceA)

	fmt.Printf("sliceB => %v\n", sliceB)
	printLen(sliceB)
	printCap(sliceB)
	sliceB = sliceB[:cap(sliceB)]
	fmt.Printf("sliceB => %v\n", sliceB)
	printLen(sliceB)
	printCap(sliceB)

}

func goRoutineLoopVariable() {
	values := []string{"this", "is", "a", "slice", "of", "strings"}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
}

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

type contextKey string

func contextStuff() {
	parentCtx := context.Background()
	key := contextKey("name")
	childCtx := context.WithValue(parentCtx, key, "Jack")

	fmt.Printf("value => %q\n", childCtx.Value(key))
}

func contextGoroutineStuff() {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, 3*time.Second)
	data := make(chan string, 1)

	go func() {
		fmt.Print("working.")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		data <- "the string you waited for"
		fmt.Print("\n")
	}()

	select {
	case message := <-data:
		fmt.Printf("message => %q\n", message)
	case <-ctx.Done():
		fmt.Println("\ncontext cancelled.")
	}
}

func main() {
	// contextStuff()
	contextGoroutineStuff()
}

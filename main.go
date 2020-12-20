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
	sliceB := arr[2:3]

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
	// auto cancel(timeout) after 3 seconds
	ctx, cancel := context.WithTimeout(parentCtx, 4*time.Second)
	defer cancel()
	data := make(chan string)

	go func() {
		fmt.Print("working.")
		// send data to channel after 5 seconds
		// for i := 0; i < 5; i++ {
		// time.Sleep(1 * time.Second)
		// fmt.Print(".")
		// }
		data <- "the string you waited for"
		fmt.Println("\ndone")
	}()

	// wait for first signal on either data chan or ctx.Done() chan
	select {
	case message := <-data:
		fmt.Printf("message => %q\n", message)
	case <-ctx.Done():
		fmt.Println("\ncontext cancelled.")
	}
}

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

func pingPongGoroutine() {

	type Ball struct{ hits int }

	table := make(chan *Ball)

	player := func(name string, table chan *Ball) {
		// fmt.Printf("%q now playing\n", name)
		for {
			ball := <-table
			ball.hits++
			fmt.Println(name, ball.hits)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		}
	}
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab  the ball
}
func main() {
	pingPongGoroutine()
}

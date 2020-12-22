package main

import (
	"context"
	"fmt"
	"time"
)

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
package main

import (
	"fmt"
	"time"
)

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
	go player("pong", table)
	go player("ping", table)

	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)
	// wait for last play then grab ball
	<-table // game over; grab  the ball
}

package main

import "fmt"

func directedChannels() {

	ping := func(pings chan<- string, msg string) {
		pings <- fmt.Sprintf("ping: [%s]", msg)
	}

	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- fmt.Sprintf("pong: [%s]", msg)
	}
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

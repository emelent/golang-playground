package main

import (
	"fmt"
	"strings"
)

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

func pipelineSample() {

	generate := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}

			close(out)
		}()

		return out
	}

	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}

			close(out)
		}()

		return out
	}

	numbersChan := generate(2, 3, 4, 5)
	squaredChan := square(numbersChan)

	printChan := func(name string, in <-chan int) {
		s := strings.Builder{}
		s.WriteString("[ ")
		for n := range in {
			s.WriteString(fmt.Sprintf("%d ", n))
		}
		s.WriteString("]")

		fmt.Printf("%s => %s\n", name, s.String())
	}

	// this interferes because chan values can only be read once,
	// so reading this, while square goroutine is  still running
	// creates race condition, and will lead to unexpected results

	// printChan("Numbers", numbersChan)

	printChan("Result", squaredChan)

}

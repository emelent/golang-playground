package main

import (
	"fmt"
	"strings"
	"sync"
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

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func pipelineSample() {

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
	squaredChan := square(square(numbersChan))

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

	in := generate(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := square(in)
	c2 := square(in)

	/*   mergeFake := func(channels ...<-chan int) <-chan int {
	 *     out := make(chan int)
	 *     var wg sync.WaitGroup
	 *     wg.Add(len(channels))
	 *     for _, c := range channels {
	 *       in := c
	 *       go func() {
	 *         for n := range in {
	 *           out <- n
	 *         }
	 *         wg.Done()
	 *       }()
	 *     }
	 *
	 *     go func() {
	 *       wg.Wait()
	 *       close(out)
	 *     }()
	 *     return out
	 *   } */

	merge := func(channels ...<-chan int) <-chan int {
		var wg sync.WaitGroup
		wg.Add(len(channels))

		out := make(chan int)

		output := func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}

		for _, c := range channels {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Printf("%d ", n) // 4 then 9, or 9 then 4
	}
	fmt.Println()

}

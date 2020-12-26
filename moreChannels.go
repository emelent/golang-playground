package main

import (
	"fmt"
	"strings"
	"sync"
)

func stringify(c <-chan interface{}) <-chan string {
	out := make(chan string)

	go func() {
		for v := range c {
			out <- fmt.Sprintf("%v", v)
		}

		close(out)
	}()

	return out
}

func fanningOut() {

	in := generate(1, 2, 3)
	var mainWg sync.WaitGroup
	mainWg.Add(4)

	printStrChan := func(c <-chan int) {
		sb := strings.Builder{}
		for n := range c {
			sb.WriteString(fmt.Sprintf("%d ", n))
		}
		fmt.Println(sb.String())
		mainWg.Done()
	}

	out1 := make(chan int)
	out2 := make(chan int)
	out3 := make(chan int)

	fanOut := func(in <-chan int, outs ...chan<- int) {

		for n := range in {
			var wg sync.WaitGroup
			wg.Add(len(outs))
			for i, c := range outs {
				out := c
				v := n
				index := i
				go func() {
					out <- v
					fmt.Printf("c[%d] <- %d\n", index, v)
					wg.Done()
				}()
			}
			wg.Wait()
		}

		for _, out := range outs {
			close(out)
		}
		mainWg.Done()

	}

	go fanOut(in, out1, out2, out3)

	go printStrChan(out1)
	go printStrChan(out2)
	go printStrChan(out3)

	mainWg.Wait()
}

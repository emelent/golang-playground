package main

import "fmt"

func goRoutineLoopVariable() {
	values := []string{"this", "is", "a", "slice", "of", "strings"}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
}
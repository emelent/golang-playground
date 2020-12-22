package main

import "fmt"

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
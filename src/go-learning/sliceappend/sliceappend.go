package main

import "fmt"

func main() {
	var list []int
	printSlice(list)

	list = append(list, 0) // Unlike python is a function
	printSlice(list)

	list = append(list, 12)
	printSlice(list)


	list = append(list,2, 3, 4)
	printSlice(list)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
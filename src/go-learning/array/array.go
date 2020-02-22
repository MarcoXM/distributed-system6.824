package main

import "fmt"

func main() {
	var a [2]string 
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	primesNumbers := [8]int{2,3,5,7,11,13,17,19}
	fmt.Println(primesNumbers)

	names := [4]string{
		
		"James",
		"Wade",
		"Kobe",
		"Durent",
	}

	fmt.Println(names)

	eastern := names[0:2]
	western := names[2:]
	fmt.Println(eastern,western)

	western[1] = "RIP"
	fmt.Println(eastern)
	fmt.Println(western)

}
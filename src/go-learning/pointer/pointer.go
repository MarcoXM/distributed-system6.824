package main

import "fmt"


func main () {
	i, j := 42, 2701 // define the i, j values

	p := &i  // pointer p would be 42
	fmt.Println(*p) // getting the value p point to
	*p = 21 // here not := we are not define.
	fmt.Println(i)

	p = &j
	*p = *p/37 // refine j by pointer p
	fmt.Println(j)
}

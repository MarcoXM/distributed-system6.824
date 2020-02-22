package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("This is q slice", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("This is r slice", r)

	// define data type no comma
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 5) // it is like the np.zeros(dtype = np.int)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	
}

//not interpretate line by line
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

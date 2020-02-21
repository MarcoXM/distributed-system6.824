package main

import "fmt"


type myStruct struct {
	X int
	Y int
	Z int
}

var (
	v1 = myStruct{1, 2, 3}
	v2 = myStruct{X:4, Y:5}
        v3 = myStruct{}  // 000
        p1 = &myStruct{7, 8, 9}
)
func main() {

	v := myStruct{4, 8, 12}
	v.X = 100
	fmt.Println(v.X,v.Y,v.Z)
	fmt.Println("Let us check the pointer with truct")

	p := &v // now *p is the v
	p.X = 1e10
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}

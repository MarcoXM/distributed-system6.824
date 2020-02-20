package main

import "fmt"

func swap(x, y string) (string,string) {
	return y, x
}


func split(sum int) (x, y int) {
	x = sum * 4 /9
	y = sum - x 
	return
}

var i, j int = 0, 1

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(45))
	var C, python, java = true, false, "no!"
	fmt.Println(i, j , C, python, java)

}

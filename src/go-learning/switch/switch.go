package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	t := time.Now()
        switch {
        case t.Hour() < 12:
                fmt.Println("Good Morning!!")
        case t.Hour() < 17:
                fmt.Println("Good Afternoon !!")
        default:
                fmt.Println("Time for Eating!!")
        }


}



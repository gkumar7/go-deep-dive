package main

import (
	"fmt"
)

type MyInt int

func (i MyInt) Val() {
	fmt.Printf("The value I am holding is %d\n", i)
}

func main() {
	i := 1
	// i.Val() = type int has no field or method Val, but... // HL
	MyInt(i).Val()
}

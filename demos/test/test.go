package main

import (
	"fmt"
)

type Worker struct {
	done chan bool
}

func (w *Worker) print(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	
}

func main() {
	for
}

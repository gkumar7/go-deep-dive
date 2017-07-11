package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Worker struct {
}

func main() {
	outputFileName := "/tmp/output-local.txt"

	f, err := os.Open(outputFileName)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", outputFileName, err)
	}

	for in := bufio.NewScanner(f); in.Scan(); {
		fmt.Println(in.Text())
	}

}

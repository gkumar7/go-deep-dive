package main

import (
	"log"
	"os"
)

func main() {
	filename := "/tmp/test"
	_, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", filename, err)
	}
}

package main

import (
	"fmt"
	"reflect"
)

// interface{} matches anything (concept can be compared to java Object declaration)
func printElem(name string, elem interface{}) { // HL
	fmt.Printf("%s has type %s\n", name, reflect.TypeOf(elem))
}

func main() {
	var explicitType int = 1.0
	inferredType := 1.0

	printElem("explicitType", explicitType)
	printElem("inferredType", inferredType)
}

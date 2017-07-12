package main

import "log"

type MischeviousGopher struct {
	name string
}

func (m MischeviousGopher) Error() string {
	return m.name
}

func ReturnsError() error {
	return MischeviousGopher{name: "George"}
}

func main() {
	err := ReturnsError()
	if err != nil {
		log.Fatal(err)
	}
}

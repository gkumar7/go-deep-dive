package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// START WRITER OMIT

type OutputWriter struct {
	outputFile string
}

func (o OutputWriter) Write(data []byte) (int, error) {
	f, err := os.OpenFile(o.outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", o.outputFile, err)
	}

	defer f.Close() // HL

	return f.Write(data)
}

// END WRITER OMIT

// START RANDOM GENERATOR OMIT

var _ Generator = (*RandomGenerator)(nil)

type Generator interface {
	Generate(int) string
}

type RandomGenerator struct {
	OutputWriter
}

func (rg *RandomGenerator) Generate(count int) string {
	var result bytes.Buffer
	rand.Seed(time.Now().UTC().UnixNano())

	delim := "\n"
	for i := 0; i < count; i++ {
		result.WriteString(delim)
		result.WriteString(strconv.Itoa(rand.Int()))
		delim = " "
	}

	return result.String()
}

// END RANDOM GENERATOR OMIT

// START SEQUENCE GENERATOR OMIT

var _ Generator = (*SequenceGenerator)(nil)

type SequenceGenerator struct {
	OutputWriter
}

func (sg *SequenceGenerator) Generate(count int) string {
	difference := 2
	delim := "\n"
	var result bytes.Buffer
	nextInt := 0
	for i := 0; i < count; i++ {
		result.WriteString(delim)
		result.WriteString(strconv.Itoa(nextInt))
		nextInt += difference
		delim = " "
	}
	return result.String()

}

// END SEQUENCE GENERATOR OMIT

// START RUN OMIT
func generateAndWrite(g Generator, count int) {
	switch v := g.(type) {
	case *RandomGenerator:
		io.WriteString(v, v.Generate(count)) // HL
	case *SequenceGenerator:
		io.WriteString(v, v.Generate(count)) // HL
	}

}

func runCmd(cmd string) (result string) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("Unable to execute command %s: %v", cmd, err)
	}
	result = string(out)
	return
}

// END RUN OMIT

// START MAIN OMIT

func main() {
	outputFileName := "/tmp/output-local.txt"
	outputWriter := OutputWriter{outputFile: outputFileName}
	generators := []Generator{
		&RandomGenerator{OutputWriter: outputWriter},
		&SequenceGenerator{OutputWriter: outputWriter},
	}

	// Formatted string
	runCmd(fmt.Sprintf("rm %s", outputFileName))

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		randomGenerator := rand.Intn(2)
		generateAndWrite(generators[randomGenerator], 10)
	}

	fmt.Println(runCmd("cat " + outputFileName))
}

// END MAIN OMIT

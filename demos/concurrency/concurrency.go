package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// START LEADER1 OMIT
type Leader struct {
	opts LeaderOpts
	pool chan string
}

type LeaderOpts struct {
	filename    string
	WorkerCount int
}

func (l *Leader) startWorkers() []Worker {
	workers := make([]Worker, 0)
	for i := 0; i < l.opts.WorkerCount; i++ {
		worker := Worker{
			id:      i,
			pool:    &l.pool,
			results: make(chan map[string]int),
		}

		workers = append(workers, worker)
		go worker.process()
	}
	return workers
}

// END LEADER1 OMIT

func printResults(results map[string]int) {
	fmt.Printf("\nTotal: \n")
	for k, v := range results {
		fmt.Printf("%s: %d\n", k, v)
	}
}

// START LEADER2 OMIT
func (l *Leader) process() {
	f, err := os.Open(l.opts.filename)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", l.opts.filename, err)
	}

	// Start worker routines which will read from pool
	workers := l.startWorkers()

	for in := bufio.NewScanner(f); in.Scan(); {
		l.pool <- in.Text()
	}
	close(l.pool)

	// Aggregate counts
	total := make(map[string]int)
	for _, worker := range workers {
		for counts := range worker.results {
			for k, v := range counts {
				total[k] += v
			}
		}
	}

	printResults(total)
}

// END LEADER2 OMIT

// START WORKER OMIT

type Worker struct {
	id      int
	pool    *chan string
	results chan map[string]int
}

func (w *Worker) process() {
	defer close(w.results)

	interim := make(map[string]int)

	for line := range *w.pool {
		fmt.Printf("Worker %d read line %s\n", w.id, line)
		vals := strings.Split(line, " ")
		for _, val := range vals {
			interim[val]++
		}
	}

	w.results <- interim
}

// END WORKER OMIT

// START MAIN OMIT
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: ./concurrency <filename> <workerCount>\n")
	}

	WorkerCount, _ := strconv.Atoi(os.Args[2]) // HL
	opts := LeaderOpts{
		filename:    os.Args[1],
		WorkerCount: WorkerCount,
	}

	l := Leader{
		opts: opts,
		pool: make(chan string),
	}

	l.process()
}

// END MAIN OMIT

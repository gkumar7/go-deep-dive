package main

// START DEF OMIT
import (
	"fmt"
	"sync"
)

const (
	RequestCount = iota
	ErrorCount
)

var (
	metricsContainer = MetricsContainer{metrics: make(map[int]int)}
	finished         = make(chan bool)
)

type MetricsContainer struct {
	sync.Mutex
	metrics map[int]int
}

// END DEF OMIT

// START MAIN OMIT

type Counter struct{}

func (c *Counter) inc(metricName int, count int) {
	metricsContainer.Lock()
	defer metricsContainer.Unlock()
	metricsContainer.metrics[metricName] += count
	finished <- true
}

func main() {
	numCounters := 5
	for i := 0; i < numCounters; i++ {
		go (&Counter{}).inc(RequestCount, 1)
	}

	// Wait for counters to complete
	for i := 0; i < numCounters; i++ {
		<-finished
	}

	// Print resulting metric
	fmt.Printf("RequestCounter has val %d\n", metricsContainer.metrics[RequestCount])
}

// END MAIN OMIT

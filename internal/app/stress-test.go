package stress_test

import (
	"fmt"
	"sync"
	"time"

	http_client "github.com/gabrielsc1998/go-stress-test/internal/infra"
)

type StressTest struct {
	Url         string
	Requests    int
	Concurrency int
}

type StressTestResult struct {
	TotalTime     float64
	TotalRequests int
	StatusCodes   map[int]int
	Errors        []error
}

var mutex = &sync.Mutex{}

func New(url string, requests int, concurrency int) *StressTest {
	return &StressTest{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

func (s *StressTest) Run() {
	var testInit = time.Now()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "- Running stress test...")

	var wg sync.WaitGroup

	var result *StressTestResult = &StressTestResult{
		StatusCodes: make(map[int]int),
	}

	for i := 0; i < s.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.run(result)
		}()
	}

	wg.Wait()

	result.TotalTime = time.Since(testInit).Seconds()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "- Stress test finished!")
	s.printResult(result)
}

func (s *StressTest) printResult(result *StressTestResult) {
	fmt.Println("\n---------- Results ----------")
	fmt.Printf("\n- Total time: %.2fs\n", result.TotalTime)
	fmt.Printf("- Total requests: %d\n", result.TotalRequests)
	fmt.Println("- Status codes:")
	for code, count := range result.StatusCodes {
		fmt.Printf(" - %d: %d\n", code, count)
	}
	if len(result.Errors) != 0 {
		fmt.Println("- Errors:")
		for _, err := range result.Errors {
			fmt.Printf(" - %s\n", err.Error())
		}
	}
	fmt.Println("\n-----------------------------")
}

func (s *StressTest) run(result *StressTestResult) {
	for i := 0; i < s.Requests/s.Concurrency; i++ {
		resp, err := http_client.Get(s.Url)

		mutex.Lock()

		result.TotalRequests++
		if err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.StatusCodes[resp.StatusCode]++
		}

		mutex.Unlock()
	}
}
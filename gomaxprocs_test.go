package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutine)

	group.Wait()
}

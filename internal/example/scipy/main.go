package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

func main() {
	st := time.Now()
	defer func() {
		fmt.Println("done:", time.Since(st))
	}()

	wg := sync.WaitGroup{}
	const n = 10000

	targets := []float64{1, 0}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			trainX := []int{1, 2, 3, 4, 5}
			trainY := []int{1, 2, 3, 4, 5}
			arr := GetPopt(trainX, trainY)
			if len(arr) != len(targets) {
				os.Exit(-1)
			}

			for i := 0; i < len(arr); i++ {
				if math.Abs(arr[i]-targets[i]) > 1e-8 {
					os.Exit(-1)
				}
			}
		}()
	}
	wg.Wait()
}

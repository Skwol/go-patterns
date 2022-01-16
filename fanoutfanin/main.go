package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	rand := func() int64 { return int64(rand.Intn(50000000)) }
	randIntStream := repeatFn(ctx, rand)

	start := time.Now()
	withoutFanOutFanIn(ctx, randIntStream)
	fmt.Printf("Search without FanOutFanIn took: %v\n", time.Since(start))

	start = time.Now()
	withFanOutFanIn(ctx, randIntStream)
	fmt.Printf("Search with FanOutFanIn took: %v\n", time.Since(start))
}

func repeatFn(ctx context.Context, fn func() int64) <-chan int64 {
	valueStream := make(chan int64)
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-ctx.Done():
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func withoutFanOutFanIn(ctx context.Context, randIntStream <-chan int64) {
	fmt.Println("Primes:")
	for prime := range take(ctx, primeFinder(ctx, randIntStream), 10) {
		fmt.Printf("%d\t", prime)
	}
	fmt.Println()
}

func take(ctx context.Context, valueStream <-chan int64, num int) <-chan int64 {
	takeStream := make(chan int64)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func primeFinder(ctx context.Context, randInt <-chan int64) <-chan int64 {
	// naive implementation of verifying if number is prime
	result := make(chan int64)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-ctx.Done():
				return
			case val := <-randInt:
				if val <= 1 {
					result <- val
					continue
				}
				for i := int64(2); i < val; i++ {
					if val%i == 0 {
						continue LOOP
					}
				}
				result <- val
			}
		}
	}()
	return result
}

func withFanOutFanIn(ctx context.Context, randIntStream <-chan int64) {
	// fan out is spitting operations on several routines
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan int64, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(ctx, randIntStream)
	}
	// operations were splitted
	for prime := range take(ctx, fanIn(ctx, finders...), 10) {
		fmt.Printf("%d\t", prime)
	}
	fmt.Println()
}

func fanIn(ctx context.Context, channels ...<-chan int64) <-chan int64 {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int64)
	multiplex := func(c <-chan int64) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}
	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

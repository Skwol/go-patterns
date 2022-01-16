package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	repeat := RepeatGenerator(ctx, 1, 2, 3, 4)
	for v := range Take(ctx, repeat, 5) {
		fmt.Printf("value from repeat %v\n", v)
	}

	rand := func() interface{} { return rand.Intn(50000000) }
	repeat = RepeatFnGenerator(ctx, rand)
	for v := range Take(ctx, repeat, 5) {
		fmt.Printf("value from repeat function %v\n", v)
	}
}

// RepeatGenerator recieves some values and infinitely sends them to the output channel
func RepeatGenerator(ctx context.Context, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// RepeatGenerator recieves some functions and infinitely runs it. All results sent to channel
func RepeatFnGenerator(ctx context.Context, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
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

// Take usefull function to request some number of values from infinite generator (channel)
func Take(ctx context.Context, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
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

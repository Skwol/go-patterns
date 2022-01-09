package main

import "fmt"

type computer struct {
	isRunning bool
}

func (t *computer) on() {
	t.isRunning = true
	fmt.Println("Turning computer on")
}

func (t *computer) off() {
	t.isRunning = false
	fmt.Println("Turning computer off")
}

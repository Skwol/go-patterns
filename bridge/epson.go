package main

import "fmt"

type lightning struct{}

func (ls *lightning) startCharging() {
	fmt.Println("Start charging with a lightning")
}

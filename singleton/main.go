package main

import "time"

func main() {
	for i := 0; i < 5; i++ {
		go getInstance()
	}
	time.Sleep(time.Second)
}

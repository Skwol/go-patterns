package main

import "fmt"

type transportlayer struct {
	next layer
}

func (l *transportlayer) execute(r *request) {
	if r.transportLayerDone {
		fmt.Println("Transport layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("transport layer request processing")
	r.transportLayerDone = true
	l.next.execute(r)
}

func (l *transportlayer) setNext(next layer) {
	l.next = next
}

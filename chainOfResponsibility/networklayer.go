package main

import "fmt"

type networklayer struct {
	next layer
}

func (l *networklayer) execute(r *request) {
	if r.networkLayerDone {
		fmt.Println("Network layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("network layer request processing")
	r.networkLayerDone = true
	l.next.execute(r)
}

func (l *networklayer) setNext(next layer) {
	l.next = next
}

package main

import "fmt"

type physicallayer struct {
	next layer
}

func (l *physicallayer) execute(r *request) {
	if r.physicalLayerDone {
		fmt.Println("Physical layer already processed")
		return
	}
	fmt.Println("physical layer request processing")
	r.physicalLayerDone = true
}

func (l *physicallayer) setNext(next layer) {
	l.next = next
}

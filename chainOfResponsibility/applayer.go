package main

import "fmt"

type applicationlayer struct {
	next layer
}

func (l *applicationlayer) execute(r *request) {
	if l == nil {
	}
	if r.appLayerDone {
		fmt.Println("Application layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("application layer request processing")
	r.appLayerDone = true
	l.next.execute(r)
}

func (l *applicationlayer) setNext(next layer) {
	l.next = next
}

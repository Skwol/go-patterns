package main

import "fmt"

type presentationlayer struct {
	next layer
}

func (l *presentationlayer) execute(r *request) {
	if r.presLayerDone {
		fmt.Println("Presentation layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("presentation layer request processing")
	r.presLayerDone = true
	l.next.execute(r)
}

func (l *presentationlayer) setNext(next layer) {
	l.next = next
}

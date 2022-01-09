package main

import "fmt"

type datalayer struct {
	next layer
}

func (l *datalayer) execute(r *request) {
	if r.dataLayerDone {
		fmt.Println("Data layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("data layer request processing")
	r.dataLayerDone = true
	l.next.execute(r)
}

func (l *datalayer) setNext(next layer) {
	l.next = next
}

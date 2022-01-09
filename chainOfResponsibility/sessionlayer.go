package main

import "fmt"

type sessionlayer struct {
	next layer
}

func (l *sessionlayer) execute(r *request) {
	if r.sessionLayerDone {
		fmt.Println("Session layer already processed")
		l.next.execute(r)
		return
	}
	fmt.Println("session layer request processing")
	r.sessionLayerDone = true
	l.next.execute(r)
}

func (l *sessionlayer) setNext(next layer) {
	l.next = next
}

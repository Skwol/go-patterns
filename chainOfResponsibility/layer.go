package main

type layer interface {
	execute(*request)
	setNext(layer)
}

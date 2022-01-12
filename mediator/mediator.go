package main

type mediator interface {
	canArrive(plane) bool
	notifyAboutDeparture()
}

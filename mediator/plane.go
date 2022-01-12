package main

import "fmt"

type plane interface {
	arrive()
	depart()
	permitArrival()
}

type passengerPlane struct {
	mediator mediator
}

func (p passengerPlane) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("PassengerPlane: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerPlane: Arrived")
}

func (p passengerPlane) depart() {
	fmt.Println("PassengerPlane: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p passengerPlane) permitArrival() {
	fmt.Println("PassengerPlane: Arrival permitted, arriving")
	p.arrive()
}

type cargoPlane struct {
	mediator mediator
}

func (p cargoPlane) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("CargoPlane: Arrival blocked, waiting")
		return
	}
	fmt.Println("CargoPlane: arrived")
}

func (p cargoPlane) depart() {
	fmt.Println("CargoPlane: departing")
	p.mediator.notifyAboutDeparture()
}

func (p cargoPlane) permitArrival() {
	fmt.Println("CargoPlane: arrival permitted")
	p.arrive()
}

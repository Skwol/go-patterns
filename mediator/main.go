package main

func main() {
	controller := newAirTrafficController()

	passengerPlane := &passengerPlane{mediator: controller}
	cargoPlane := &cargoPlane{mediator: controller}

	passengerPlane.arrive()
	cargoPlane.arrive()
	passengerPlane.depart()
}

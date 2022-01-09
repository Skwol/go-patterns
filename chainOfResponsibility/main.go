package main

func main() {
	physicalLayer := &physicallayer{}

	dataLayer := &datalayer{}
	dataLayer.setNext(physicalLayer)

	networkLayer := &networklayer{}
	networkLayer.setNext(dataLayer)

	transportLayer := &transportlayer{}
	transportLayer.setNext(networkLayer)

	sessionLayer := &sessionlayer{}
	sessionLayer.setNext(transportLayer)

	presLayer := &presentationlayer{}
	presLayer.setNext(sessionLayer)

	appLayer := &applicationlayer{}
	appLayer.setNext(presLayer)

	r := &request{name: "test"}
	appLayer.execute(r)
}

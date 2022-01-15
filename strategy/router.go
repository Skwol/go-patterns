package main

import "fmt"

type router interface {
	calculateRoute(n *navigator) error
}

type onFootRouter struct{}

func (r onFootRouter) calculateRoute(n *navigator) error {
	fmt.Println("calculating route on foot")
	return nil
}

type byBusRouter struct{}

func (r byBusRouter) calculateRoute(n *navigator) error {
	fmt.Println("calculating route by bus")
	return nil
}

type byCarRouter struct{}

func (r byCarRouter) calculateRoute(n *navigator) error {
	fmt.Println("calculating route by car")
	return nil
}

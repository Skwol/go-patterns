package main

import "fmt"

type sportswearFactory interface {
	makeShoe() sneakers
	makeShirt() tshirt
}

func getSportsFactory(brand string) (sportswearFactory, error) {
	if brand == "reebok" {
		return &reebok{}, nil
	}

	if brand == "newbalance" {
		return &newbalance{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

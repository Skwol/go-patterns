package main

import "fmt"

func getAnimal(animalType string) (animal, error) {
	switch animalType {
	case "dog":
		return newDog(), nil
	case "cat":
		return newCat(), nil
	default:
		return nil, fmt.Errorf("unexpected animal type")
	}
}

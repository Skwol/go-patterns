package main

import "fmt"

type europeanHumidifier struct{}

func (w *europeanHumidifier) plugIn() {
	fmt.Println("plug in european socket.")
}

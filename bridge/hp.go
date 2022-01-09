package main

import "fmt"

type usbc struct{}

func (u *usbc) startCharging() {
	fmt.Println("Start charging with a usb type c")
}

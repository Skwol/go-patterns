package main

import "fmt"

type client struct{}

func (c *client) plugInAmericanSocket(com humidifier) {
	fmt.Println("Client plugs in humidifier.")
	com.plugInAmericanSocket()
}

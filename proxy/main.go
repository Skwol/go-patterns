package main

import "fmt"

const (
	stateURL  = "/state"
	createURL = "/create"
)

func main() {
	nginxServer := newProxy()

	c, b := nginxServer.handleRequest(stateURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", stateURL, c, b)

	c, b = nginxServer.handleRequest(stateURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", stateURL, c, b)

	c, b = nginxServer.handleRequest(stateURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", stateURL, c, b)

	c, b = nginxServer.handleRequest(createURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", stateURL, c, b)

	c, b = nginxServer.handleRequest(createURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", stateURL, c, b)
}

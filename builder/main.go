package main

import "fmt"

func main() {
	tcpBuilder := getBuilder("tcp")
	unixBuilder := getBuilder("unix")

	director := newDirector(tcpBuilder)
	tcpResponse := director.buildResponse()

	fmt.Printf("tcp response status code: %d\n", tcpResponse.statusCode)
	fmt.Printf("tcp response content type: %s\n", tcpResponse.contentType)
	fmt.Printf("tcp response protocol: %s\n", tcpResponse.prototcol)

	director.setBuilder(unixBuilder)
	socketResponse := director.buildResponse()

	fmt.Printf("socket response status code: %d\n", socketResponse.statusCode)
	fmt.Printf("socket response content type: %s\n", socketResponse.contentType)
	fmt.Printf("socket response protocol: %s\n", socketResponse.prototcol)
}

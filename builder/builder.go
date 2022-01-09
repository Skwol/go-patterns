package main

type builder interface {
	setContentType()
	setStatusCode()
	setProtocol()
	getHouse() response
}

func getBuilder(builderType string) builder {
	if builderType == "tcp" {
		return &tcpBuilder{}
	}

	if builderType == "unix" {
		return &socketBuilder{}
	}
	return nil
}

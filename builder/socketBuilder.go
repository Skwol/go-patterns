package main

type socketBuilder struct {
	contentType string
	statusCode  int
	prototcol   string
}

func newIglooBuilder() *socketBuilder {
	return &socketBuilder{}
}

func (b *socketBuilder) setContentType() {
	b.contentType = "json"
}

func (b *socketBuilder) setStatusCode() {
	b.statusCode = 200
}

func (b *socketBuilder) setProtocol() {
	b.prototcol = "unix"
}

func (b *socketBuilder) getHouse() response {
	return response{
		statusCode:  b.statusCode,
		contentType: b.contentType,
		prototcol:   b.prototcol,
	}
}

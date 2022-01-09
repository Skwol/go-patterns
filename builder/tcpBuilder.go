package main

type tcpBuilder struct {
	contentType string
	statusCode  int
	protocol    string
}

func newNormalBuilder() *tcpBuilder {
	return &tcpBuilder{}
}

func (b *tcpBuilder) setContentType() {
	b.contentType = "xml"
}

func (b *tcpBuilder) setStatusCode() {
	b.statusCode = 201
}

func (b *tcpBuilder) setProtocol() {
	b.protocol = "tcp"
}

func (b *tcpBuilder) getHouse() response {
	return response{
		statusCode:  b.statusCode,
		contentType: b.contentType,
		prototcol:   b.protocol,
	}
}

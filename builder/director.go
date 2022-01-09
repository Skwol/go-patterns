package main

type director struct {
	builder builder
}

func newDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b builder) {
	d.builder = b
}

func (d *director) buildResponse() response {
	d.builder.setStatusCode()
	d.builder.setContentType()
	d.builder.setProtocol()
	return d.builder.getHouse()
}

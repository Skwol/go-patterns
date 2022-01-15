package main

type carContext struct {
	engineOn  state
	engineOff state

	currentState state
	speed        int
}

func newCar() *carContext {
	var c carContext
	c.engineOff = &engineOffState{c: &c}
	c.engineOn = &engineOnState{c: &c}
	c.setState(c.engineOff)
	return &c
}

func (c *carContext) setState(s state) {
	c.currentState = s
}

func (c *carContext) startEngine() error {
	return c.currentState.startEngine()
}

func (c *carContext) stopEngine() error {
	return c.currentState.stopEngine()
}

func (c *carContext) setSpeed(speed int) error {
	return c.currentState.setSpeed(speed)
}

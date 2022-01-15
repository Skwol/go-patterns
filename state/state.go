package main

import "fmt"

type state interface {
	startEngine() error
	stopEngine() error
	setSpeed(int) error
}

type engineOnState struct {
	c *carContext
}

func (o *engineOnState) startEngine() error {
	return fmt.Errorf("engine alrady started")
}

func (o *engineOnState) stopEngine() error {
	if o.c.speed > 0 {
		return fmt.Errorf("need to slow down first")
	}
	fmt.Println("stopping car engine")
	o.c.setState(o.c.engineOff)
	return nil
}

func (o *engineOnState) setSpeed(speed int) error {
	fmt.Println("setting car speed to", speed)
	o.c.speed = speed
	return nil
}

type engineOffState struct {
	c *carContext
}

func (o *engineOffState) startEngine() error {
	fmt.Println("starting engine")
	o.c.setState(o.c.engineOn)
	return nil
}

func (o *engineOffState) stopEngine() error {
	return fmt.Errorf("engine is off")
}

func (o *engineOffState) setSpeed(speed int) error {
	return fmt.Errorf("engine is off")
}

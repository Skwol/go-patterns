package main

import "fmt"

type unit interface {
	move(direction)
	stop()
	attack()
}

type gameUnit struct {
	unit unit
}

type direction string

const (
	upDirection    = direction("up")
	downDirection  = direction("down")
	leftDirection  = direction("left")
	rightDirection = direction("right")
)

var directionReverse = map[direction]direction{
	upDirection:    downDirection,
	downDirection:  upDirection,
	leftDirection:  rightDirection,
	rightDirection: leftDirection,
}

func (d direction) reverse() direction {
	return directionReverse[d]
}

func (u *gameUnit) hitAndRunTemplateMethod(d direction) error {
	u.unit.move(d)
	u.unit.stop()
	u.unit.attack()
	u.unit.move(d.reverse())
	u.unit.stop()
	return nil
}

type cavarlymen struct {
	gameUnit
}

func (c *cavarlymen) move(d direction) {
	fmt.Printf("cavalry gallops %s\n", d)
}

func (c *cavarlymen) stop() {
	fmt.Printf("cavalry gallops without stopping\n")
}

func (c *cavarlymen) attack() {
	fmt.Printf("cavalry crashes into enemy formation\n")
}

type archer struct {
	gameUnit
}

func (a *archer) move(d direction) {
	fmt.Printf("archer moves %s\n", d)
}

func (a *archer) stop() {
	fmt.Printf("archer stops a hundred paces from the enemy\n")
}

func (a *archer) attack() {
	fmt.Printf("archer fires at the enemy\n")
}

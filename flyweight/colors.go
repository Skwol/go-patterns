package main

import "fmt"

const (
	RedColor   = "r"
	GreenColor = "g"
)

// color can be common thing for thousands of figures
type color interface {
	Name() string
}

type redColor struct {
	name string
}

func (r redColor) Name() string {
	return r.name
}

type greenColor struct {
	name string
}

func (g greenColor) Name() string {
	return g.name
}

var colorFactoryInstance = &colorFactory{
	colorMap: make(map[string]color),
}

type colorFactory struct {
	colorMap map[string]color
}

func (c *colorFactory) GetColorByName(colorName string) (color, error) {
	if v, ok := c.colorMap[colorName]; ok {
		return v, nil
	}

	switch colorName {
	case RedColor:
		c.colorMap[colorName] = &redColor{name: "red"}
		return c.colorMap[colorName], nil
	case GreenColor:
		c.colorMap[colorName] = &greenColor{name: "green"}
		return c.colorMap[colorName], nil
	default:
		return nil, fmt.Errorf("Wrong dress type passed")
	}
}

func getColorFactory() *colorFactory {
	return colorFactoryInstance
}

package main

type figure struct {
	color color
	area  float64
}

func newFigure(figureType, colorStr string) *figure {
	color, _ := getColorFactory().GetColorByName(colorStr)
	return &figure{
		color: color,
		area:  10,
	}
}

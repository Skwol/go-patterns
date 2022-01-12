package main

import "fmt"

func main() {
	screen := &screen{}

	screen.addCircle(RedColor)
	screen.addCircle(GreenColor)

	screen.addRectangle(GreenColor)
	screen.addRectangle(RedColor)

	colorFactory := getColorFactory()

	total := 0
	for colorName, color := range colorFactory.colorMap {
		fmt.Printf("color key: %s\t color name: %s\n", colorName, color.Name())
		total++
	}
	fmt.Printf("total number of color instances %d\n", total)
}

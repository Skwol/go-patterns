package main

type screen struct {
	rectangles []*figure
	circles    []*figure
}

func (s *screen) addRectangle(color string) {
	rect := newFigure("rect", color)
	s.rectangles = append(s.rectangles, rect)
	return
}

func (s *screen) addCircle(color string) {
	circ := newFigure("circ", color)
	s.circles = append(s.circles, circ)
	return
}

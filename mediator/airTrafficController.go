package main

type airTrafficController struct {
	isAirstripFree bool
	planeQueue     []plane
}

func newAirTrafficController() *airTrafficController {
	return &airTrafficController{
		isAirstripFree: true,
	}
}

func (c *airTrafficController) canArrive(p plane) bool {
	if c.isAirstripFree {
		c.isAirstripFree = false
		return true
	}
	c.planeQueue = append(c.planeQueue, p)
	return false
}

func (s *airTrafficController) notifyAboutDeparture() {
	if !s.isAirstripFree {
		s.isAirstripFree = true
	}
	if len(s.planeQueue) > 0 {
		s.planeQueue[0].permitArrival()
		s.planeQueue = s.planeQueue[1:]
	}
}

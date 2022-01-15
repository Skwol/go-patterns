package main

type navigator struct {
	// some extrimely important data here
	router router
}

func newNavigator() navigator {
	return navigator{}
}

func (n *navigator) setRouter(r router) {
	n.router = r
}

func (n *navigator) calculateRoute(pointA, pointB string) error {
	// some magic with coordinates here
	return n.router.calculateRoute(n)
}

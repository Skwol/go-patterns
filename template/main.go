package main

func main() {
	cavalery := &cavarlymen{}
	unit := gameUnit{cavalery}
	unit.hitAndRunTemplateMethod(upDirection)

	arch := &archer{}
	unit = gameUnit{arch}
	unit.hitAndRunTemplateMethod(leftDirection)
}

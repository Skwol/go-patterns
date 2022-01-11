package main

import "fmt"

func main() {
	snapshots := &snapshotHandler{states: make([]*stateHolder, 0)}
	originator := &originator{state: "1"}

	fmt.Printf("status: %s\n", originator.getState())
	snapshots.addSnapshot(originator.createSnapshot())

	originator.setState("2")
	fmt.Printf("status: %s\n", originator.getState())
	snapshots.addSnapshot(originator.createSnapshot())

	originator.setState("3")
	fmt.Printf("status: %s\n", originator.getState())
	snapshots.addSnapshot(originator.createSnapshot())

	originator.restoreSnapshot(snapshots.getSnapshot(1))
	fmt.Printf("status: %s\n", originator.getState())

	originator.restoreSnapshot(snapshots.getSnapshot(0))
	fmt.Printf("status: %s\n", originator.getState())
}

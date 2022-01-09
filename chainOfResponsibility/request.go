package main

type request struct {
	name               string
	appLayerDone       bool
	presLayerDone      bool
	sessionLayerDone   bool
	transportLayerDone bool
	networkLayerDone   bool
	dataLayerDone      bool
	physicalLayerDone  bool
}

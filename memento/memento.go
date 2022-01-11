package main

type snapshotHandler struct {
	states []*stateHolder
}

func (c *snapshotHandler) addSnapshot(m *stateHolder) {
	c.states = append(c.states, m)
}

func (c snapshotHandler) getSnapshot(index int) *stateHolder {
	return c.states[index]
}

type stateHolder struct {
	state string
}

func (s *stateHolder) getSavedState() string {
	return s.state
}

type originator struct {
	state string
}

func (o *originator) createSnapshot() *stateHolder {
	return &stateHolder{state: o.state}
}

func (o *originator) restoreSnapshot(s *stateHolder) {
	o.state = s.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

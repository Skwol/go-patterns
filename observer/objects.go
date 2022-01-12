package main

import "fmt"

type sku struct {
	name      string
	available bool
	observers []observer
}

func newItem(name string) *sku {
	return &sku{
		name: name,
	}
}

func (i *sku) updateAvailability() {
	fmt.Printf("%s is now in available\n", i.name)
	i.available = true
	i.notifyAll()
}

func (i *sku) register(o observer) {
	i.observers = append(i.observers, o)
}

func (s *sku) deregister(o observer) {
	for i, observer := range s.observers {
		if o.ID() == observer.ID() {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (i *sku) notifyAll() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("notify to customer %s for %s availability\n", c.id, itemName)
}

func (c *customer) ID() string {
	return c.id
}

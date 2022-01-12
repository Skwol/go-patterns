package main

type observer interface {
	update(string)
	ID() string
}

type subject interface {
	register(observer)
	deregister(observer)
	notify()
}

package main

type lister interface {
	iterator(*cat) iterator
}

type iterator interface {
	next() bool
}

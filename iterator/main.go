package main

import "fmt"

func main() {
	catOne := &cat{
		name: "jimbo",
		age:  3,
	}
	catTwo := &cat{
		name: "jambo",
		age:  4,
	}

	cats := &catList{
		cats: []*cat{catOne, catTwo},
	}

	var c cat
	iter := cats.iterator(&c)
	for iter.next() {
		fmt.Printf("cat is %+v\n", c)
	}
}

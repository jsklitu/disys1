package main

type Fork struct {
	free bool
	used int

	in  chan<- int
	out chan<- int
}

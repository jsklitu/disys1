package main

import (
	f "disys1/forks"
	p "disys1/philosophers"
)

func main() {
	//evt bare lav 5 af hver og så 10 kanaler
	//når fork of ohil oprettes skal de så navngives med noget smart

	ch1, ch2 := make(chan string), make(chan string)
	go f.Fork(ch1, ch2)
	go p.Phil(ch2, ch1)
	for {
		//eat forever
	}
}

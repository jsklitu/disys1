package main

import (
	f "disys1/forks"
	p "disys1/philosophers"
	"fmt"
	"sync"
)

func main() {
	var waitGroup = sync.WaitGroup{}
	waitGroup.Add(1)

	mainIn := make(chan string)
	mainOut := make(chan string)

	var chanArray [20]chan string

	for i := range chanArray {
		chanArray[i] = make(chan string)
	}

	f1 := f.Fork{ID: 1, LeftIn: chanArray[3], LeftOut: chanArray[2], RightIn: chanArray[0], RightOut: chanArray[1], MainIn: mainIn, MainOut: mainOut}
	f2 := f.Fork{ID: 2, LeftIn: chanArray[7], LeftOut: chanArray[6], RightIn: chanArray[4], RightOut: chanArray[5], MainIn: mainIn, MainOut: mainOut}
	f3 := f.Fork{ID: 3, LeftIn: chanArray[11], LeftOut: chanArray[10], RightIn: chanArray[8], RightOut: chanArray[9], MainIn: mainIn, MainOut: mainOut}
	f4 := f.Fork{ID: 4, LeftIn: chanArray[15], LeftOut: chanArray[14], RightIn: chanArray[12], RightOut: chanArray[13], MainIn: mainIn, MainOut: mainOut}
	f5 := f.Fork{ID: 5, LeftIn: chanArray[19], LeftOut: chanArray[18], RightIn: chanArray[16], RightOut: chanArray[17], MainIn: mainIn, MainOut: mainOut}
	forkArray := []*f.Fork{&f1, &f2, &f3, &f4, &f5}

	p1 := p.Philosopher{Name: "A", LeftIn: chanArray[18], LeftOut: chanArray[19], RightIn: chanArray[1], RightOut: chanArray[0], MainIn: mainIn, MainOut: mainOut}
	p2 := p.Philosopher{Name: "B", LeftIn: chanArray[5], LeftOut: chanArray[4], RightIn: chanArray[2], RightOut: chanArray[3], MainIn: mainIn, MainOut: mainOut}
	p3 := p.Philosopher{Name: "C", LeftIn: chanArray[6], LeftOut: chanArray[7], RightIn: chanArray[9], RightOut: chanArray[8], MainIn: mainIn, MainOut: mainOut}
	p4 := p.Philosopher{Name: "D", LeftIn: chanArray[13], LeftOut: chanArray[12], RightIn: chanArray[10], RightOut: chanArray[11], MainIn: mainIn, MainOut: mainOut}
	p5 := p.Philosopher{Name: "E", LeftIn: chanArray[17], LeftOut: chanArray[16], RightIn: chanArray[14], RightOut: chanArray[15], MainIn: mainIn, MainOut: mainOut}
	philArray := []*p.Philosopher{&p1, &p2, &p3, &p4, &p5}

	go mainChannels(mainIn, mainOut)

	for i := range forkArray {
		go f.ForkFunc(*forkArray[i])
	}

	for i := range philArray {
		go p.Phil(*philArray[i])
	}

	waitGroup.Wait()
}

func mainChannels(mainIn, mainOut chan string) {
	for {

		mainIn <- "pls give me info"

		fmt.Println(<-mainOut)
	}

}

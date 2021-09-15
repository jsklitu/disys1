package forks

import (
	"fmt"
	"strconv"
)

type Fork struct {
	ID int

	LeftIn  chan string
	LeftOut chan string

	RightIn  chan string
	RightOut chan string

	MainIn  chan string
	MainOut chan string
}

func ForkFunc(f Fork) {
	counter := 0
	for {

		select {

		case <-f.LeftIn:
			{
				f.LeftOut <- "You got me b"
				counter++
				fmt.Println("Fork " + strconv.Itoa(f.ID) + " is NOT available")
				<-f.LeftIn
				fmt.Println("Fork " + strconv.Itoa(f.ID) + " is available")
			}

		case <-f.RightIn:
			{
				f.RightOut <- "You got me b"
				counter++
				fmt.Println("Fork " + strconv.Itoa(f.ID) + " is NOT available")
				<-f.RightIn
				fmt.Println("Fork " + strconv.Itoa(f.ID) + " is available")
			}
		}

		select {
		case <-f.MainIn:
			{
				f.MainOut <- "Fork " + strconv.Itoa(f.ID) + " has been used " + strconv.Itoa(counter) + " times"
			}

		}
	}
}

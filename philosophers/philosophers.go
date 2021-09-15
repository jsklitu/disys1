package philosophers

import (
	"fmt"
	"strconv"
	"time"
)

type Philosopher struct {
	Name string

	LeftIn  chan string
	LeftOut chan string

	RightIn  chan string
	RightOut chan string

	MainIn  chan string
	MainOut chan string
}

func Phil(p Philosopher) {
	counter := 0
	for {
		p.LeftOut <- "i wanna eat pls"
		<-p.LeftIn //get left fork

		p.RightOut <- "are you also available i wanna EAATTTTT!!"
		<-p.RightIn //get right fork

		fmt.Println("Philosopher " + p.Name + " is now eating")
		counter++

		time.Sleep(time.Second)

		p.LeftOut <- "u available now"
		p.RightOut <- "u2"

		fmt.Println("Philosopher " + p.Name + " is now thinking")

		select {
		case <-p.MainIn:
			{
				p.MainOut <- "Philosopher " + p.Name + " has now eaten " + strconv.Itoa(counter) + " times"
			}

		}

	}
}

package main

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go fork(ch1, ch2)
	go phil(ch2, ch1)
	for {
		//eat forever
	}
}

func fork(chIN, chOUT) {
	for {
		<-chIN
		//besked om at samlpe op

		chOUT <- "spis, yes"

		//hvornår er jeg lagt igen?
		<-chIN
	}
}

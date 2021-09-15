package forks

func Fork(chIN, chOUT chan string) {
	for {
		<-chIN
		//besked om at samlpe op

		chOUT <- "spis, yes"

		//hvornår er jeg lagt igen?
		<-chIN
	}
}

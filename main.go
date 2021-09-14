package main

func main() {
	//we need to be able to ask with certain int values
	// 0 = how many times have you been used/eaten
	// -1 = are you being used/are you eating

}

//channelse. start med channels, og så bag efter kig på at optælle

func main(){
	ch1,ch2 :_ make(chan string), make (chan string)
	go fork (ch1, ch2)
	go phil(ch2, ch1)
	for{
		//eat forever
	}
}


func phil(chIN, chOut){
	counter:= 0
	for{
	chOUT <- i wanna eat	
	<-chIN
	counter++
	chOUT <- im done

	}
}

func fork(chIN, chOUT){
	for{
		<-chIN
		//besked om at samlpe op

		chOUT <- "spis, yes"

		//hvornår er jeg lagt igen?
		<-chIN	
	}
}
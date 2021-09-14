package main

import(
	"fmt"
)

type Philosopher struct{
	//channel in - spiser du lige nu? hvor mange gange har du spist?
	in chan <- int

	//channel out - jeg har spist så mange gange, og jeg tænker
	out chan <- int
	
	//number of times eaten
	var timesEaten int

	//eating (true) or thinking (false)
	var eating bool
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

//channel fra phil til gaffel. tre tjeks i alt. er du ledig? ja. nu er jeg færidg. og så til begge sider
//ch1:=make(chan string)
	//<- chIN - jeg vil gerne modtage. når den får en besked kan den skrive til chOUT "tag mig", som phil ser
	//når phil så er færdig får den endnu en besked på chIN "nu er jeg done"
//med unbuffered står den stille og venter på der kommer en besked. først når der kommer en går den
//videre. og slef skal der være en goroutine der skal læse inputtet



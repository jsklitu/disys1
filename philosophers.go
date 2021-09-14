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


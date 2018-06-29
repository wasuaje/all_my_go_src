package main


import (
	"fmt"
	"goroutines/common"
)



func main() {
	var c chan string = make(chan string)

	go goroutines.Pinger(c)
	go goroutines.Ponger(c)
	go goroutines.Panger(c)
	go goroutines.Printer(c)

	var input string
	fmt.Scanln( &input)
}


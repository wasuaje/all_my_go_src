package goroutines


import (
	"fmt"
	"time"
)

func Pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Panger(c chan string) {
	for i := 0; ; i++ {
		c <- "pang"
	}
}

func Printer(c chan string) {
	for {
		fmt.Println(<- c, len(c))
		time.Sleep(time.Second * 1)
	}
}



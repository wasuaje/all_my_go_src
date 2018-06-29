
package main

import (
 "fmt"
 "time"
"log"
)

var(
	currentTurn = 1
	totalTurns =  5
)

func pomodoroTurn(chanPomodoro chan bool) {
	tellBeginTurn()
	time.Sleep(time.Second * 10) // Replace *time.Minute* by *time.Second* for quicker testing
	tellEndTurn()
	chanPomodoro <- true
}

func pomodoroBreak(chanBreak chan bool) {
	tellBeginSmallBreak()
	time.Sleep(time.Second * 5)
	tellEndSmallBreak()
	chanBreak <- true
}

func pomodoroLongBreak(chanLongBreak chan bool) {
	tellBeginLongBreak()
	time.Sleep(time.Second * 30)
	tellEndLongBreak()
	chanLongBreak <- true
}

func tellBeginTurn() {
	log.Print("Pomodoro round begins")
}

func tellEndTurn() {
	log.Print("Round ended")
}

func tellBeginSmallBreak() {
	log.Print("Have a small break!")
}

func tellEndSmallBreak() {
	log.Print("This is the end of the small break. Let's go back to work!")
}

func tellBeginLongBreak() {
	log.Print("Have a long break! You deserved it!")
}

func tellEndLongBreak() {
	log.Print("This is the end of the long break. Let's go back to work!")
}

func pomodoroService(chanPomodoro, chanBreak, chanLongBreak, chanDone chan bool) {
	fmt.Println("Pomodoro service started\n")
	for {
		select {

		case endTurn := <-chanPomodoro:
			_ = endTurn
			if currentTurn >= totalTurns {
				go pomodoroLongBreak(chanLongBreak)
				currentTurn = 1
			} else {
				currentTurn += 1
				go pomodoroBreak(chanBreak)
			}

		case endSmallBreak := <-chanBreak:
			_ = endSmallBreak
			go pomodoroTurn(chanPomodoro)

		case endLongBreak := <-chanLongBreak:
			_ = endLongBreak
			input := askAnotherSession()
			for input != "Y" && input != "N" {
				input = askAnotherSession()
			}
			if input == "Y" {
				go pomodoroTurn(chanPomodoro)
			} else {
				chanDone <- true
			}

		}
	}
}

func askAnotherSession() string {
	fmt.Println("Ready for another pomodoro session? (Y/N)")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	turn := make(chan bool)
	smallBreak := make(chan bool)
	longBreak := make(chan bool)
	done := make(chan bool)

	go pomodoroTurn(turn)
	go pomodoroService(turn, smallBreak, longBreak, done)

	<-done
}

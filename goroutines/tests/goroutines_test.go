package goroutines

import ("testing"
"goroutines/common"

)


func TestPinger(t *testing.T) {
	t.Log("Testing Pinger... (expecting: 'ping')")
	var c chan string = make(chan string)
	go goroutines.Pinger(c)
	if val := <-c; val != "ping"{
		t.Errorf("Expected 'ping', but it was '%v' instead.", val)
	}

}

func TestPonger(t *testing.T) {
	t.Log("Testing Ponger... (expecting: 'pong')")
	var c chan string = make(chan string)
	go goroutines.Ponger(c)
	if val := <-c; val != "pong"{
		t.Errorf("Expected 'pong', but it was '%v' instead.", val)
	}

}


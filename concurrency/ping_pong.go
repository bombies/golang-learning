package concurrency

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, numPings)
	func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}

// don't touch below this line

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := range numPings {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func testPong(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func TestPingPong() {
	testPong(4)
	testPong(3)
	testPong(2)
}

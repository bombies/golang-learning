package concurrency

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for range numDBs {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := range numDBs {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func addEmailsToQueue(emails []string) chan string {
	bufferedChannel := make(chan string, len(emails))

	for _, email := range emails {
		bufferedChannel <- email
	}

	return bufferedChannel
}

func countReports(numSentCh chan int) int {
	count := 0
	for v, ok := <-numSentCh; ok; v, ok = <-numSentCh {
		count += v
	}
	return count
}

func sendReports(numBatches int, ch chan int) {
	for i := range numBatches {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func concurrentFib(n int) []int {
	intChan := make(chan int)
	fib := make([]int, 0, n)

	go fibonacci(n, intChan)

	for num := range intChan {
		fib = append(fib, num)
	}

	return fib
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

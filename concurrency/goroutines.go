package concurrency

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func TestGoroutines(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

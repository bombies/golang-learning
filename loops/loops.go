package loops

import "log"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := range numMessages {
		totalCost += 1 + float64(i)/100
	}
	return totalCost
}

func maxMessages(thresh int) int {
	totalCost, messageCount := 100, 0

	for ; totalCost <= thresh; messageCount++ {
		totalCost += 100 + messageCount
	}

	return messageCount
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for actualCostInPennies < balance {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func FizzBuzz() {
	for i := range 100 {
		switch {
		case i%3 == 0:
			log.Println("fizz")
		case i%5 == 0:
			log.Println("buzz")
		default:
			log.Println("fizzbuzz")
		}
	}
}

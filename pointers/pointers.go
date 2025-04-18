package pointers

import (
	"strings"

	"github.com/samber/lo"
)

func removeProfanity(message *string) {
	profanityMap := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}

	*message = lo.Reduce(lo.Entries(profanityMap), func(agg string, item lo.Entry[string, string], index int) string {
		return strings.ReplaceAll(agg, item.Key, item.Value)
	}, *message)
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(analytics *Analytics, msg Message) {
	switch {
	case msg.Success:
		analytics.MessagesSucceeded += 1
	default:
		analytics.MessagesFailed += 1
	}

	analytics.MessagesTotal += 1
}

/*
If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic)
that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying
to dereference it.
*/
func removeProfanity2(message *string) {
	if message == nil {
		return
	}

	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

/*
A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common
than value receivers. However, methods with pointer receivers don't require that a
pointer is used to call the method. The pointer will automatically be derived from the value.
*/
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

package structs

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	userIsValid := func(u user) bool {
		return u.name != "" && u.number > 0
	}

	return userIsValid(mToSend.sender) && userIsValid(mToSend.recipient) && mToSend.message != ""
}

func TestStructs() {
	mToSend := messageToSend{
		message:   "you have an appointment tomorrow",
		sender:    user{name: "Brenda Halafax", number: 16545550987},
		recipient: user{name: "Sally Sue", number: 19035558973},
	}

	fmt.Printf("%s can send to %s: %t\n", mToSend.sender.name, mToSend.recipient.name, canSendMessage(mToSend))
}

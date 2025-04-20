package enums

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}

	err = a.track(em.status)
	if err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}

	return nil
}

// This is a sumulation of union types in typescript
type sendingChannel string

const (
	Email sendingChannel = "email"
	SMS   sendingChannel = "sms"
	Phone sendingChannel = "phone"
)

// We can use the special `iota` keyword to mimic enum behaviour
type sendingChannel2 int

const (
	Email2 sendingChannel2 = iota
	SMS2
	Phone2
)

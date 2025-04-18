package interfaces

import "fmt"

// Mutliple interface implementation
type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

func (e email) cost() int {
	switch {
	case e.isSubscribed:
		return 2 * len(e.body)
	default:
		return 5 * len(e.body)
	}
}

func (e email) format() string {
	return fmt.Sprintf("'%s' | %s", e.body, func() string {
		if e.isSubscribed {
			return "Subscribed"
		}
		return "Not Subscribed"
	}())
}

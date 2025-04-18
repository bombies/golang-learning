package interfaces

type new_expense interface {
	newCost() float64
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) newCost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) newCost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) newCost() float64 {
	return 0.0
}

func getExpenseReport(e new_expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.newCost()
	case sms:
		return v.toPhoneNumber, v.newCost()
	default:
		return "", float64(v.newCost())
	}
}

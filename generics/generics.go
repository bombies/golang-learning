package generics

import (
	"errors"
	"fmt"
	"time"
)

func getLast[T any](s []T) T {
	var last T

	if len(s) == 0 {
		return last
	}

	return s[len(s)-1]
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) (hist []T, newBalance float64, err error) {
	newItemCost := newItem.GetCost()

	if newItemCost > balance {
		return nil, 0.0, errors.New("insufficient funds")
	}

	hist = append(oldItems, newItem)
	newBalance = balance - newItemCost
	return hist, newBalance, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

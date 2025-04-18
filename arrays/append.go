package arrays

import "github.com/samber/lo"

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// Ideally I was supposed to the use append() builtin function to manually add the costs
	// to a new arrow but that's lame. For future reference just know it exists.
	// https://pkg.go.dev/builtin#append
	
	return lo.Map(
		lo.Filter(costs, func(item cost, _ int) bool {
			return item.day == day
		}),
		func(item cost, _ int) float64 {
			return item.value
		},
	)
}

package arrays

import "github.com/samber/lo"

func sum(nums ...int) int {
	return lo.Reduce(nums, func(agg, item, idx int) int {
		return agg + item
	}, 0)
}

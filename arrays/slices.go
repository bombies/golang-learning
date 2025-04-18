package arrays

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case "pro":
		return messages[:], nil
	case "free":
		return messages[:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for idx, msg := range messages {
		costs[idx] = float64(len(msg)) * 0.01
	}

	return costs
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0, rows)

	for x := range rows {
		row := make([]int, 0, cols)
		for y := range cols {
			row = append(row, x*y)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

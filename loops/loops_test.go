package loops

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test_connections_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		numMessages int
		expected    float64
	}
	runCases := []testCase{
		{10, 10.45},
		{20, 21.9},
	}

	submitCases := append(runCases, []testCase{
		{0, 0.0},
		{1, 1.0},
		{5, 5.10},
		{30, 34.35},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := bulkSend(test.numMessages)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.numMessages, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.numMessages, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test2(t *testing.T) {
	type testCase struct {
		thresh   int
		expected int
	}
	runCases := []testCase{
		{103, 1},
		{205, 2},
		{1000, 9},
	}

	submitCases := append(runCases, []testCase{
		{100, 1},
		{3000, 26},
		{4000, 34},
		{5000, 41},
		{0, 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := maxMessages(test.thresh)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.thresh, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.thresh, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test3(t *testing.T) {
	type testCase struct {
		costMultiplier   float64
		maxCostInPennies int
		expected         int
	}

	runCases := []testCase{
		{1.1, 5, 4},
		{1.3, 10, 5},
		{1.35, 25, 7},
	}

	submitCases := append(runCases, []testCase{
		{1.2, 1, 1},
		{1.2, 15, 7},
		{1.3, 20, 7},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMaxMessagesToSend(test.costMultiplier, test.maxCostInPennies)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_connections_challenge(t *testing.T) {
	type testCase struct {
		groupSize int
		expected  int
	}

	runCases := []testCase{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
	}

	submitCases := append(runCases, []testCase{
		{0, 0},
		{10, 45},
		{100, 4950},
		{1000, 499500},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result := countConnections(test.groupSize)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Fail
`, test.groupSize, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Pass
`, test.groupSize, test.expected, result)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

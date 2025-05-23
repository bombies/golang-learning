package generics

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
}

func test1(t *testing.T) {
	type testCase struct {
		input    any
		expected any
	}
	runCases := []testCase{
		{[]int{}, 0},
		{[]bool{true, false, true, true, false}, false},
	}

	submitCases := append(runCases, []testCase{
		{[]int{1, 2, 3, 4}, 4},
		{[]string{"a", "b", "c", "d"}, "d"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0

	for _, test := range testCases {
		switch v := test.input.(type) {
		case []int:
			if output := getLast(v); output != test.expected {
				t.Errorf(`
---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`
---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []string:
			if output := getLast(v); output != test.expected {
				t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []bool:
			if output := getLast(v); output != test.expected {
				t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}
}

func test2(t *testing.T) {
	type testCase struct {
		newItem           lineItem
		oldItems          []lineItem
		balance           float64
		expected          []lineItem
		expectedBalance   float64
		expectedErrString string
	}

	runCases := []testCase{
		{
			newItem: subscription{
				userEmail: "geralt@rivia.com",
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				interval:  "yearly",
			},
			oldItems: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
			},
			balance: 1000.00,
			expected: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
				subscription{
					userEmail: "geralt@rivia.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "yearly",
				},
			},
			expectedBalance:   750.00,
			expectedErrString: "",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			newItem: subscription{
				userEmail: "geralt@rivia.com",
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				interval:  "yearly",
			},
			oldItems: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
			},
			balance:           200.00,
			expected:          nil,
			expectedBalance:   0.0,
			expectedErrString: "insufficient funds",
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		oldItems := slices.Clone(test.oldItems)
		newItems, newBalance, err := chargeForLineItem(
			test.newItem,
			test.oldItems,
			test.balance,
		)
		if (err != nil && err.Error() != test.expectedErrString) ||
			(err == nil && test.expectedErrString != "") ||
			!slices.Equal(newItems, test.expected) ||
			newBalance != test.expectedBalance {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  newItem:  %v
  oldItems:
%v
  balance:  %v
  expected items:
%v
  expected balance: %v
  expected error:   %v
  actual items:
%v
  actual balance: %v
  actual error:   %v
`,
				test.newItem,
				sliceWithBullets(oldItems),
				test.balance,
				sliceWithBullets(test.expected),
				test.expectedBalance,
				test.expectedErrString,
				sliceWithBullets(newItems),
				newBalance,
				err,
			)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  newItem:  %v
  oldItems:
%v
  balance:  %v
  expected items:
%v
  expected balance: %v
  expected error:   %v
  actual items:
%v
  actual balance: %v
  actual error:   %v
`,
				test.newItem,
				sliceWithBullets(oldItems),
				test.balance,
				sliceWithBullets(test.expected),
				test.expectedBalance,
				test.expectedErrString,
				sliceWithBullets(newItems),
				newBalance,
				err,
			)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

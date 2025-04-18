package arrays

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test4(t)
	test5(t)
	test6(t)
	test_message_filter_challenge(t)
	test_valid_password_challenge(t)
	test_tagger_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		messages         []string
		expectedMessages [3]string
		expectedCosts    [3]int
	}

	runCases := []testCase{
		{
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]int{46, 92, 119},
		},
		{
			[]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]int{27, 49, 61},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]int{21, 42, 59},
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
		actualMessages, actualCosts := getMessageWithRetries(test.messages[0], test.messages[1], test.messages[2])
		if actualMessages[0] != test.expectedMessages[0] ||
			actualMessages[1] != test.expectedMessages[1] ||
			actualMessages[2] != test.expectedMessages[2] ||
			actualCosts[0] != test.expectedCosts[0] ||
			actualCosts[1] != test.expectedCosts[1] ||
			actualCosts[2] != test.expectedCosts[2] {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
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
		plan             string
		messages         [3]string
		expectedMessages []string
		expectedErr      string
	}
	runCases := []testCase{
		{
			planFree,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{"Hello sir/madam can I interest you in a yacht?", "Please I'll even give you an Amazon gift card?"},
			"",
		},
		{
			planPro,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			"",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			planFree,
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			[]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
			},
			"",
		},
		{
			planPro,
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			[]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			"",
		},
		{
			"invalid plan",
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			nil,
			"unsupported plan",
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
		actualMessages, err := getMessageWithRetriesForPlan(test.plan, test.messages)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if !slices.Equal(actualMessages, test.expectedMessages) || errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Fail
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Pass
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
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
		messages    []string
		expected    []float64
		expectedCap int
	}

	runCases := []testCase{
		{
			[]string{"Welcome to the movies!", "Enjoy your popcorn!"},
			[]float64{0.22, 0.19},
			2,
		},
		{
			[]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"},
			[]float64{0.31, 0.15, 0.1, 0.09},
			4,
		},
	}

	submitCases := append(runCases, []testCase{
		{[]string{}, []float64{}, 0},
		{[]string{""}, []float64{0}, 1},
		{[]string{"Hello", "Hi", "Hey"}, []float64{0.05, 0.02, 0.03}, 3},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMessageCosts(test.messages)
		if !slicesEqual(output, test.expected) || cap(output) != test.expectedCap {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test4(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}

	runCases := []testCase{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	submitCases := append(runCases, []testCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := sum(test.nums...)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Fail
`, sliceWithBullets(test.nums), test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Pass
`, sliceWithBullets(test.nums), test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test5(t *testing.T) {
	type testCase struct {
		costs    []cost
		day      int
		expected []float64
	}

	runCases := []testCase{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{5, 2.5},
				{2, 3.6},
				{1, 2.7},
				{1, 3.3},
			},
			day: 1,
			expected: []float64{
				2.0,
				3.1,
				2.7,
				3.3,
			},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{2, 2.5},
				{3, 3.1},
				{3, 2.6},
				{4, 3.34},
			},
			day: 4,
			expected: []float64{
				3.34,
			},
		},
		{
			costs: []cost{
				{0, 1.0},
				{10, 2.0},
				{3, 3.1},
				{2, 2.5},
				{1, 3.6},
				{2, 2.7},
				{4, 56.34},
				{13, 2.34},
				{28, 1.34},
				{25, 2.34},
				{30, 4.34},
			},
			day:      5,
			expected: []float64{},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		output := getDayCosts(test.costs, test.day)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
%v
Expecting:
%v
Actual:
%v
Fail
`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %v
Expecting:
%v
Actual:
%v
Pass
`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test6(t *testing.T) {
	type testCase struct {
		rows, cols int
		expected   [][]int
	}

	runCases := []testCase{
		{3, 3, [][]int{
			{0, 0, 0},
			{0, 1, 2},
			{0, 2, 4},
		}},
		{4, 4, [][]int{
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 2, 4, 6},
			{0, 3, 6, 9},
		}},
	}

	submitCases := append(runCases, []testCase{
		{5, 7, [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 2, 3, 4, 5, 6},
			{0, 2, 4, 6, 8, 10, 12},
			{0, 3, 6, 9, 12, 15, 18},
			{0, 4, 8, 12, 16, 20, 24},
		}},
		{0, 0, [][]int{}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := createMatrix(test.rows, test.cols)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed: %v x %v matrix
Expecting:
%v
Actual:
%v
Fail
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed: %v x %v matrix
Expecting:
%v
Actual:
%v
Pass
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_message_filter_challenge(t *testing.T) {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}
	type testCase struct {
		filterType    string
		expectedCount int
		expectedType  string
	}

	runCases := []testCase{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
	}

	submitCases := append(runCases, []testCase{
		{"media", 2, "media"},
		{"text", 2, "text"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			filtered := filterMessages(messages, test.filterType)
			if len(filtered) != test.expectedCount {
				failCount++
				t.Errorf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Fail
`, i+1, test.filterType, test.expectedCount, len(filtered))
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Pass
`, i+1, test.filterType, test.expectedCount, len(filtered))
			}

			for _, m := range filtered {
				if m.Type() != test.expectedType {
					failCount++
					t.Errorf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Fail
`, i+1, test.expectedType, m.Type())
				} else {
					passCount++
					fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Pass
`, i+1, test.expectedType, m.Type())
				}
			}
		})
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_valid_password_challenge(t *testing.T) {
	type testCase struct {
		password string
		isValid  bool
	}

	runCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
	}

	submitCases := append(runCases, []testCase{
		{"VeryLongPassword1", false},
		{"Short", false},
		{"1234short", false},
		{"Short5", true},
		{"P4ssword", true},
		{"AA0Z9", true},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			result := isValidPassword(test.password)
			if result != test.isValid {
				failCount++
				t.Errorf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Fail
`, test.password, test.isValid, result)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Pass
`, test.password, test.isValid, result)
			}
		})
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_tagger_challenge(t *testing.T) {
	type testCase struct {
		messages []sms
		expected [][]string
	}

	runCases := []testCase{
		{
			messages: []sms{{id: "001", content: "Urgent, please respond!"}, {id: "002", content: "Big sale on all items!"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
		},
		{
			messages: []sms{{id: "003", content: "Enjoy your day"}},
			expected: [][]string{{}},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			messages: []sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}},
			expected: [][]string{{"Urgent", "Promo"}},
		},
		{
			messages: []sms{{id: "005", content: "i nEEd URgEnt help, my FROZEN FLAME was used"}, {id: "006", content: "wAnt to saLE 200x heavy leather"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
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
		actual := tagMessages(test.messages, tagger)
		if len(actual) != len(test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed for length of returned sms slice
Expecting: %v
Actual:    %v
Fail
`, len(test.expected), len(actual))
			continue
		}

		for i, msg := range actual {
			if !reflect.DeepEqual(msg.tags, test.expected[i]) {
				failCount++
				t.Errorf(`---------------------------------
Test Failed for message ID %s
Expecting: %v
Actual:    %v
Fail
`, msg.id, test.expected[i], msg.tags)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for message ID %s
Expecting: %v
Actual:    %v
Pass
`, msg.id, test.expected[i], msg.tags)
			}
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

func slicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func formatMatrix(matrix [][]int) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%v\n", row)
	}
	return result
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

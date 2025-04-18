package interfaces

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test4(t)
	test_message_formatter_challenge(t)
	test_process_notifications_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		msg          message
		expectedText string
		expectedCost int
	}

	runCases := []testCase{
		{birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"},
			"Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			168,
		},
		{sendingReport{"First Report", 10},
			`Your "First Report" report is ready. You've sent 10 messages.`,
			183,
		},
	}

	submitCases := append(runCases, []testCase{
		{birthdayMessage{time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC), "Bill Deer"},
			"Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z",
			171,
		},
		{sendingReport{"Second Report", 20},
			`Your "Second Report" report is ready. You've sent 20 messages.`,
			186,
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
		text, cost := sendMessage(test.msg)
		if text != test.expectedText || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.msg, test.expectedText, test.expectedCost, text, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msg, test.expectedText, test.expectedCost, text, cost)
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
		emp      employee
		expected int
	}

	runCases := []testCase{
		{fullTime{name: "Bob", salary: 7300}, 7300},
		{contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}, 856304},
	}

	submitCases := append(runCases, []testCase{
		{fullTime{name: "Alice", salary: 10000}, 10000},
		{contractor{name: "John", hourlyPay: 1000, hoursPerYear: 1000}, 1000000},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		salary := test.emp.getSalary()
		if salary != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Fail
`, test.emp, test.expected, salary)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Pass
`, test.emp, test.expected, salary)
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
		body           string
		isSubscribed   bool
		expectedCost   int
		expectedFormat string
	}

	runCases := []testCase{
		{"hello there", true, 22, "'hello there' | Subscribed"},
		{"general kenobi", false, 70, "'general kenobi' | Not Subscribed"},
	}

	submitCases := append(runCases, []testCase{
		{"i hate sand", true, 22, "'i hate sand' | Subscribed"},
		{"it's coarse and rough and irritating", false, 180, "'it's coarse and rough and irritating' | Not Subscribed"},
		{"and it gets everywhere", true, 44, "'and it gets everywhere' | Subscribed"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		e := email{
			body:         test.body,
			isSubscribed: test.isSubscribed,
		}
		cost := e.cost()
		format := e.format()
		if format != test.expectedFormat || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
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
		expense      new_expense
		expectedTo   string
		expectedCost float64
	}

	runCases := []testCase{
		{
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
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
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_message_formatter_challenge(t *testing.T) {
	type testCase struct {
		formatter Formatter
		expected  string
	}

	runCases := []testCase{
		{PlainText{Message: Message{message: "Hello, World!"}}, "Hello, World!"},
		{Bold{Message: Message{message: "Bold Message"}}, "**Bold Message**"},
		{Code{Message: Message{message: "Code Message"}}, "`Code Message`"},
	}

	submitCases := append(runCases, []testCase{
		{Code{Message: Message{message: ""}}, "``"},
		{Bold{Message: Message{message: ""}}, "****"},
		{PlainText{Message: Message{message: ""}}, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		testName := "Test Case " + strconv.Itoa(i+1)
		t.Run(testName, func(t *testing.T) {
			formattedMessage := SendMessage(test.formatter)
			if formattedMessage != test.expected {
				failCount++
				t.Errorf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, testName, test.formatter, test.expected, formattedMessage)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, testName, test.formatter, test.expected, formattedMessage)
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

func test_process_notifications_challenge(t *testing.T) {
	type testCase struct {
		notification       notification
		expectedID         string
		expectedImportance int
	}

	runCases := []testCase{
		{
			directMessage{senderUsername: "Kaladin", messageContent: "Life before death", priorityLevel: 10, isUrgent: true},
			"Kaladin",
			50,
		},
		{
			groupMessage{groupName: "Bridge 4", messageContent: "Soups ready!", priorityLevel: 2},
			"Bridge 4",
			2,
		},
		{
			systemAlert{alertCode: "ALERT001", messageContent: "THIS IS NOT A TEST HIGH STORM COMING SOON"},
			"ALERT001",
			100,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			directMessage{senderUsername: "Shallan", messageContent: "I am that I am.", priorityLevel: 5, isUrgent: false},
			"Shallan",
			5,
		},
		{
			groupMessage{groupName: "Knights Radiant", messageContent: "For the greater good.", priorityLevel: 10},
			"Knights Radiant",
			10,
		},
		{
			directMessage{senderUsername: "Adolin", messageContent: "Duels are my favorite.", priorityLevel: 3, isUrgent: true},
			"Adolin",
			50,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for i, test := range testCases {
		t.Run("TestProcessNotification_"+strconv.Itoa(i+1), func(t *testing.T) {
			id, importance := processNotification(test.notification)
			if id != test.expectedID || importance != test.expectedImportance {
				failCount++
				t.Errorf(`---------------------------------
Test Failed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Fail
`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Pass
`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
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

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

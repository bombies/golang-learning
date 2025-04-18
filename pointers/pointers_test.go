package pointers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test4(t)
	test_update_balance_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		messageIn string
		expected  string
	}

	runCases := []testCase{
		{
			"English, motherfubber, do you speak it?",
			"English, mother****er, do you speak it?",
		},
		{
			"Oh man I've seen some crazy ass shiz in my time...",
			"Oh man I've seen some crazy ass **** in my time...",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			"Does he look like a witch?",
			"Does he look like a *****?",
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
		original := test.messageIn
		removeProfanity(&test.messageIn)
		if test.messageIn != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
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
		initialAnalytics Analytics
		newMessage       Message
		expected         Analytics
	}

	runCases := []testCase{
		{
			initialAnalytics: Analytics{MessagesTotal: 0, MessagesFailed: 0, MessagesSucceeded: 0},
			newMessage:       Message{Recipient: "mickey", Success: true},
			expected:         Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "minnie", Success: false},
			expected:         Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			initialAnalytics: Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "goofy", Success: false},
			expected:         Analytics{MessagesTotal: 3, MessagesFailed: 2, MessagesSucceeded: 1},
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
		a := test.initialAnalytics
		getMessageText(&a, test.newMessage)
		if a != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  Initial Analytics:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  New Message:
    Recipient=%s, Success=%v
  Expected:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  Actual:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
`, test.initialAnalytics.MessagesTotal, test.initialAnalytics.MessagesFailed, test.initialAnalytics.MessagesSucceeded,
				test.newMessage.Recipient, test.newMessage.Success,
				test.expected.MessagesTotal, test.expected.MessagesFailed, test.expected.MessagesSucceeded,
				a.MessagesTotal, a.MessagesFailed, a.MessagesSucceeded)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  Initial Analytics:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  New Message:
    Recipient=%s, Success=%v
  Expected:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  Actual:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
`, test.initialAnalytics.MessagesTotal, test.initialAnalytics.MessagesFailed, test.initialAnalytics.MessagesSucceeded,
				test.newMessage.Recipient, test.newMessage.Success,
				test.expected.MessagesTotal, test.expected.MessagesFailed, test.expected.MessagesSucceeded,
				a.MessagesTotal, a.MessagesFailed, a.MessagesSucceeded)
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
		messageIn *string
		expected  *string
	}
	s1 := "English, motherfubber, do you speak it?"
	s2 := "English, mother****er, do you speak it?"
	s3 := "Does he look like a witch?"
	s4 := "Does he look like a *****?"

	runCases := []testCase{
		{
			&s1,
			&s2,
		},
		{
			nil,
			nil,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			&s3,
			&s4,
		},
		{
			nil,
			nil,
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
		var original *string
		if test.messageIn != nil {
			originalVal := *test.messageIn
			original = &originalVal
		}
		removeProfanity2(test.messageIn)
		if test.messageIn != nil &&
			test.expected != nil &&
			original != nil &&
			*test.messageIn != *test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`, *original, *test.expected, *test.messageIn)
		} else if (test.messageIn == nil || test.expected == nil) &&
			test.messageIn != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
		} else if test.messageIn == nil && test.expected == nil {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`, *original, *test.expected, *test.messageIn)
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
		e          email
		newMessage string
		expected   string
	}

	runCases := []testCase{
		{
			email{
				message:     "My name is Lt. Aldo Raine and I'm putting together a special team, and I need me eight soldiers.",
				fromAddress: "lt.aldo.raine@mailio.com",
				toAddress:   "army@mailio.com",
			},
			"You just say bingo.",
			"You just say bingo.",
		},
		{
			email{
				message:     "Now, if one were to determine what attribute the German people share with a beast, it would be the cunning and the predatory instinct of a hawk.",
				fromAddress: "col.hans.landa@mailio.com",
				toAddress:   "lapadite@mailio.com",
			},
			"What a tremendously hostile world that a rat must endure.",
			"What a tremendously hostile world that a rat must endure.",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email{
				message:     "Nazi ain't got no humanity. They're the foot soldiers of a Jew-hatin', mass murderin' maniac and they need to be dee-stroyed.",
				fromAddress: "lt.aldo.raine@mailio.com",
				toAddress:   "basterds@mailio.com",
			},
			"I think this just might be my masterpiece.",
			"I think this just might be my masterpiece.",
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
		originalMessage := test.e.message
		test.e.setMessage(test.newMessage)
		if test.e.message != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  inputs:
    * msg: %v
    * newMessage: %v
    * from: %v
    * to: %v
  expected: %v
  actual: %v
`, originalMessage, test.newMessage, test.e.fromAddress, test.e.toAddress, test.expected, test.e.message)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  inputs:
    * msg: %v
    * newMessage: %v
    * from: %v
    * to: %v
  expected: %v
  actual: %v
`, originalMessage, test.newMessage, test.e.fromAddress, test.e.toAddress, test.expected, test.e.message)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func test_update_balance_challenge(t *testing.T) {
	type testCase struct {
		name            string
		initialCustomer customer
		transaction     transaction
		expectedBalance float64
		expectError     bool
		errorMessage    string
	}

	runCases := []testCase{
		{
			name:            "Deposit operation",
			initialCustomer: customer{id: 1, balance: 100.0},
			transaction:     transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit},
			expectedBalance: 150.0,
			expectError:     false,
		},
		{
			name:            "Withdrawal operation",
			initialCustomer: customer{id: 2, balance: 200.0},
			transaction:     transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 100.0,
			expectError:     false,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			name:            "insufficient funds for withdrawal",
			initialCustomer: customer{id: 3, balance: 50.0},
			transaction:     transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 50.0,
			expectError:     true,
			errorMessage:    "insufficient funds",
		},
		{
			name:            "unknown transaction type",
			initialCustomer: customer{id: 4, balance: 100.0},
			transaction:     transaction{customerID: 4, amount: 50.0, transactionType: "unknown"},
			expectedBalance: 100.0,
			expectError:     true,
			errorMessage:    "unknown transaction type",
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
		t.Run(test.name, func(t *testing.T) {
			err := updateBalance(&test.initialCustomer, test.transaction)
			failureMessage := ""

			if (err != nil) != test.expectError {
				failureMessage += "Unexpected error presence: expected an error but didn't get one, or vice versa.\n"
			} else if err != nil && err.Error() != test.errorMessage {
				failureMessage += "Incorrect error message.\n"
			}

			if test.initialCustomer.balance != test.expectedBalance {
				failureMessage += "Balance not updated as expected.\n"
			}

			if failureMessage != "" {
				failCount++
				failureMessage = "FAIL\n" + failureMessage +
					"Transaction: " + string(test.transaction.transactionType) +
					fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
					fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, test.initialCustomer.balance)
				t.Errorf(`---------------------------------
					%s
`, failureMessage)
			} else {
				passCount++
				successMessage := "PASSED\n" +
					"Transaction: " + string(test.transaction.transactionType) +
					fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
					fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, test.initialCustomer.balance)
				fmt.Printf(`---------------------------------
%s
`, successMessage)
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

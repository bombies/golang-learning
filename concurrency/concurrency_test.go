package concurrency

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test4(t)
	test5(t)
	test6(t)
	test_process_messages_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		emails [3]email
		isOld  [3]bool
	}

	runCases := []testCase{
		{[3]email{
			{
				body: "Words are pale shadows of forgotten names. As names have power, words have power.",
				date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's like everyone tells a story about themselves inside their own head.",
				date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Bones mend. Regret stays with you forever.",
				date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, false, false}},
		{[3]email{
			{
				body: "Music is a proud, temperamental mistress.",
				date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Have you heard of that website Boot.dev?",
				date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's awesome honestly.",
				date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, true, true}},
	}

	submitCases := append(runCases, []testCase{
		{[3]email{
			{
				body: "I have stolen princesses back from sleeping barrow kings.",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I burned down the town of Trebon",
				date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I have spent the night with Felurian and left with both my sanity and my life.",
				date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, true, false}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		isOld := checkEmailAge(test.emails)
		if !slices.Equal(isOld[:], test.isOld[:]) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  input:
    * %v | %v
    * %v | %v
    * %v | %v
  expected: %v
  actual:   %v
`,
				test.emails[0].body, test.emails[0].date,
				test.emails[1].body, test.emails[1].date,
				test.emails[2].body, test.emails[2].date,
				test.isOld, isOld)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  input:
    * %v | %v
    * %v | %v
    * %v | %v
  expected: %v
  actual:   %v
`,
				test.emails[0].body, test.emails[0].date,
				test.emails[1].body, test.emails[1].date,
				test.emails[2].body, test.emails[2].date,
				test.isOld, isOld)
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
		numDBs int
	}

	runCases := []testCase{
		{1},
		{3},
		{4},
	}

	submitCases := append(runCases, []testCase{
		{0},
		{13},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0

	for _, test := range testCases {
		fmt.Printf(`---------------------------------`)
		fmt.Printf("\nTesting %v Databases...\n\n", test.numDBs)
		dbChan, count := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		for *count != test.numDBs {
			fmt.Println("...")
		}
		if len(dbChan) == 0 && *count == test.numDBs {
			passed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
PASS
`,
				test.numDBs, len(dbChan), *count)
		} else {
			failed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
FAIL
`,
				test.numDBs, len(dbChan), *count)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}
}

func test3(t *testing.T) {
	type testCase struct {
		emails   []string
		expected int
	}

	runCases := []testCase{
		{
			[]string{
				"To boldly go where no man has gone before.",
				"Live long and prosper.",
			},
			2,
		},
		{
			[]string{
				"The needs of the many outweigh the needs of the few, or the one.",
				"Change is the essential process of all existence.",
				"Resistance is futile.",
			},
			3,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{
				"It's life, Jim, but not as we know it.",
				"Infinite diversity in infinite combinations.",
				"Make it so.",
				"Engage!",
			},
			4,
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
		ch := addEmailsToQueue(test.emails)
		actual := len(ch)
		if actual != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  emails:
%v
  expected channel length: %v
  actual channel length:   %v
`,
				sliceWithBullets(test.emails),
				test.expected,
				actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  emails:
%v
  expected channel length: %v
  actual channel length:   %v
`,
				sliceWithBullets(test.emails),
				test.expected,
				actual)
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
		numBatches int
		expected   int
	}

	runCases := []testCase{
		{3, 114},
		{4, 198},
	}

	submitCases := append(runCases, []testCase{
		{0, 0},
		{1, 15},
		{6, 435},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		numSentCh := make(chan int)
		go sendReports(test.numBatches, numSentCh)
		output := countReports(numSentCh)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
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
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

func test5(t *testing.T) {
	type testCase struct {
		n        int
		expected []int
	}

	runCases := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}

	submitCases := append(runCases, []testCase{
		{0, []int{}},
		{1, []int{0}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
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
		expectedLogs []string
	}

	runCases := []testCase{
		{
			expectedLogs: []string{
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"All backups saved!",
			},
		},
	}

	submitCases := append(runCases, []testCase{}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0
	for _, test := range testCases {
		expectedLogs := test.expectedLogs

		snapshotTicker := time.Tick(800 * time.Millisecond)
		saveAfter := time.After(2800 * time.Millisecond)
		logChan := make(chan string)
		go saveBackups(snapshotTicker, saveAfter, logChan)
		actualLogs := []string{}
		for actualLog := range logChan {
			fmt.Println(actualLog)
			actualLogs = append(actualLogs, actualLog)
		}

		if !slices.Equal(expectedLogs, actualLogs) {
			t.Errorf(`---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			failed++
		} else {
			fmt.Printf(`---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			passed++
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("\n%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	}
}

func test_process_messages_challenge(t *testing.T) {
	type testCase struct {
		messages []string
		expect   []string
	}

	runCases := []testCase{
		{
			messages: []string{"Sunlit", "Man"},
			expect:   []string{"Man-processed", "Sunlit-processed"},
		},
		{
			messages: []string{"Nomad do you copy?"},
			expect:   []string{"Nomad do you copy?-processed"},
		},
		{
			messages: []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"},
			expect:   []string{"Taldain-processed", "Roshar-processed", "Sel-processed", "Nalthis-processed", "Scadriel-processed"},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			messages: []string{},
			expect:   []string{},
		},
		{
			messages: []string{"Scadriel"},
			expect:   []string{"Scadriel-processed"},
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
		fail := false
		result := processMessages(test.messages)

		if len(result) != len(test.expect) {
			fail = true
		}

		counts := make(map[string]int)
		for _, res := range result {
			counts[res]++
		}
		for _, exp := range test.expect {
			counts[exp]--
			if counts[exp] < 0 {
				fail = true
			}
		}

		if fail {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  inputs:   %v
  expected: %v
  actual:   %v
  `, test.messages, test.expect, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  inputs:   %v
  expected: %v
  actual:   %v
`, test.messages, test.expect, result)
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

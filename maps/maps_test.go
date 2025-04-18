package maps

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	test1(t)
	test2(t)
	test3(t)
	test_distinct_words_challenge(t)
	test_suggested_friends_challenge(t)
}

func test1(t *testing.T) {
	type testCase struct {
		names        []string
		phoneNumbers []int
		expected     map[string]user
		errString    string
	}
	runCases := []testCase{
		{
			[]string{"Eren", "Armin", "Mikasa"},
			[]int{14355550987, 98765550987, 18265554567},
			map[string]user{"Eren": {"Eren", 14355550987}, "Armin": {"Armin", 98765550987}, "Mikasa": {"Mikasa", 18265554567}},
			"",
		},
		{
			[]string{"Eren", "Armin"},
			[]int{14355550987, 98765550987, 18265554567},
			nil,
			"invalid sizes",
		},
	}
	submitCases := append(runCases, []testCase{
		{
			[]string{"George", "Annie", "Reiner", "Sasha"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			map[string]user{"George": {"George", 20955559812}, "Annie": {"Annie", 38385550982}, "Reiner": {"Reiner", 48265554567}, "Sasha": {"Sasha", 16045559873}},
			"",
		},
		{
			[]string{"George", "Annie", "Reiner"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			nil,
			"invalid sizes",
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
		output, err := getUserMap(test.names, test.phoneNumbers)
		if test.errString != "" && err == nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: none
`, test.names, test.phoneNumbers, test.errString)
		} else if test.errString == "" && err != nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: none
  actual err: %v
`, test.names, test.phoneNumbers, err)
		} else if test.errString != "" && err != nil && err.Error() != test.errString {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: %v
`, test.names, test.phoneNumbers, test.errString, err)
		} else if !compareMaps(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
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
		users             map[string]user2
		name              string
		expectedErrString string
		expectedUsers     map[string]user2
		expectedDeleted   bool
	}

	getMapCopy := func(m map[string]user2) map[string]user2 {
		copy := make(map[string]user2)
		for key, value := range m {
			copy[key] = value
		}
		return copy
	}

	runCases := []testCase{
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Erwin",
			"",
			map[string]user2{"Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			true,
		},
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			"Erwin",
			"",
			map[string]user2{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			false,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Eren",
			"not found",
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			false,
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
		originalUsers := getMapCopy(test.users)
		deleted, err := deleteIfNecessary(test.users, test.name)
		if test.expectedErrString != "" {
			if err == nil {
				failCount++
				t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: %v
  actual error: none
`, formatMap2(originalUsers), test.name, test.expectedErrString)
			} else if err.Error() != test.expectedErrString {
				failCount++
				t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: %v
  actual error: %v
`, formatMap2(originalUsers), test.name, test.expectedErrString, err)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed:
  users:
%v
  name: %v
  expected error: %v
  actual error: %v
`, formatMap2(originalUsers), test.name, test.expectedErrString, err)
			}
		} else if err != nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: none
  actual error: %v
`, formatMap2(originalUsers), test.name, err)
		} else if !compareMaps2(test.users, test.expectedUsers) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected users:
%v
  actual users:
%v
`, formatMap2(originalUsers), test.name, formatMap2(test.expectedUsers), formatMap2(test.users))
		} else if deleted != test.expectedDeleted {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected deleted: %v
  actual deleted: %v
`, formatMap2(originalUsers), test.name, test.expectedDeleted, deleted)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  users:
%v
  name: %v
  expected users:
%v
  actual users:
%v
  expected deleted: %v
  actual deleted: %v
`, formatMap2(originalUsers), test.name, formatMap2(test.expectedUsers), formatMap2(test.users), test.expectedDeleted, deleted)
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
		messagedUsers []string
		validUsers    map[string]int
		expected      map[string]int
	}

	runCases := []testCase{
		{
			[]string{"cersei", "jaime", "cersei"},
			map[string]int{"cersei": 0, "jaime": 0},
			map[string]int{"cersei": 2, "jaime": 1},
		},
		{
			[]string{"cersei", "tyrion", "jaime", "tyrion", "tyrion"},
			map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
			map[string]int{"cersei": 1, "jaime": 1, "tyrion": 3},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{},
			map[string]int{"tyrion": 0},
			map[string]int{"tyrion": 0},
		},
		{
			[]string{"cersei", "jaime", "tyrion"},
			map[string]int{"tywin": 0},
			map[string]int{"tywin": 0},
		},
		{
			[]string{"cersei", "cersei", "cersei", "tyrion"},
			map[string]int{"cersei": 0},
			map[string]int{"cersei": 3},
		},
		{
			[]string{"cersei", "tywin", "jaime", "cersei", "tyrion", "cersei", "jaime"},
			map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
			map[string]int{"cersei": 3, "jaime": 2, "tyrion": 1},
		},
		{
			[]string{"cersei", "cersei", "jaime", "jaime", "tywin", "cersei", "tywin", "tyrion"},
			map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
			map[string]int{"cersei": 3, "jaime": 2, "tyrion": 1},
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
		getCounts(test.messagedUsers, test.validUsers)
		if !compareMaps3(test.validUsers, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test #%v Failed:
  Messaged Users: %v
  Expected: %v
  Actual: %v
`, i, test.messagedUsers, test.expected, test.validUsers)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test #%v Passed:
  Messaged Users: %v
  Expected: %v
  Actual: %v
`, i, test.messagedUsers, test.expected, test.validUsers)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_distinct_words_challenge(t *testing.T) {
	type testCase struct {
		messages []string
		expected int
	}

	runCases := []testCase{
		{
			[]string{"WTS Arcanite Bar! Cheaper than AH", "Do you need an Arcanite Bar!"},
			10,
		},
		{
			[]string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."},
			19,
		},
		{
			[]string{"LFG UBRS", "lfg ubrs", "LFG Ubrs"},
			2,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{"Alright time's up! Let's do this.", "Leroy Jenkins!", "Damn it Leroy"},
			10,
		},
		{
			[]string{"I'm out of range", "I'm out of mana"},
			5,
		},
		{
			[]string{
				"LF9M UBRS need all",
				"LF8M UBRS need all",
				"LF7M UBRS need all",
				"LF6M UBRS need tanks and heals",
				"LF5M UBRS need tanks and heals",
				"LF4M UBRS need tanks and heals",
				"LF3M UBRS need tanks and healer",
				"LF2M UBRS need tanks",
				"LF1M UBRS need tank",
				"Group is full thanks!",
			},
			21,
		},
		{
			[]string{""},
			0,
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
		result := countDistinctWords(test.messages)
		formattedMessages := formatMessages(test.messages)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
FAIL:
Messages: %v
Expecting: %d unique words
Actual:    %d unique words
Fail
`, formattedMessages, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed!
Messages: %v
Expecting: %d unique words
Actual:    %d unique words
Pass
`, formattedMessages, test.expected, result)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func test_suggested_friends_challenge(t *testing.T) {
	friendships := map[string][]string{
		"Dalinar": {"Kaladin", "Pattern", "Shallan"},
		"Kaladin": {"Dalinar", "Syl", "Teft", "Shallan"},
		"Pattern": {"Dalinar", "Teft", "Shallan"},
		"Syl":     {"Kaladin"},
		"Teft":    {"Kaladin", "Pattern"},
		"Moash":   {},
		"Shallan": {"Pattern", "Kaladin", "Dalinar"},
	}

	type testCase struct {
		username string
		expected []string
	}

	runCases := []testCase{
		{"Dalinar", []string{"Syl", "Teft"}},
		{"Kaladin", []string{"Pattern"}},
		{"Pattern", []string{"Kaladin"}},
		{"Syl", []string{"Dalinar", "Shallan", "Teft"}},
		{"Teft", []string{"Dalinar", "Shallan", "Syl"}},
		{"Moash", nil},
	}

	submitCases := append(runCases, []testCase{
		{
			"Odium", nil,
		},
		{
			"Shallan", []string{"Syl", "Teft"},
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
		t.Run(test.username, func(t *testing.T) {
			result := findSuggestedFriends(test.username, friendships)
			sort.Strings(result)
			if !reflect.DeepEqual(result, test.expected) {
				failCount++
				t.Errorf(`---------------------------------
Test Failed for username %s:
Expecting:  %v
Actual:     %v
Fail
`, test.username, formatSlice(test.expected), formatSlice(result))
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for username %s:
Expecting:  %v
Actual:     %v
Pass
`, test.username, formatSlice(test.expected), formatSlice(result))
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

func formatSlice(slice []string) string {
	if slice == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", slice)
}

func formatMessages(messages []string) string {
	var formattedMessages []string
	for _, message := range messages {
		words := strings.Fields(message)
		formattedMessage := strings.Join(words, " ")
		formattedMessages = append(formattedMessages, fmt.Sprintf("[%s]", formattedMessage))
	}
	return strings.Join(formattedMessages, ", ")
}

func formatMap(m map[string]user) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
}

func formatMap2(m map[string]user2) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
}

func compareMaps(map1, map2 map[string]user) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

func compareMaps2(map1, map2 map[string]user2) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

func compareMaps3(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

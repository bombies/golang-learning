package maps

import (
	"errors"
	"strings"

	"github.com/samber/lo"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	users := make(map[string]user)

	for _, tuple := range lo.Zip2(names, phoneNumbers) {
		name, number := lo.Unpack2(tuple)
		users[name] = user{
			name:        name,
			phoneNumber: number,
		}
	}

	return users, nil
}

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		_, ok := validUsers[user]
		if !ok {
			continue
		}

		validUsers[user] += 1
	}
}

func countDistinctWords(messages []string) int {
	words := lo.FlatMap(messages, func(item string, _ int) []string {
		return strings.Split(strings.ToLower((item)), " ")
	})

	wordMap := make(map[string]int)

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		wordMap[word] += 1
	}

	return len(wordMap)
}

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggestedFriends := make(map[string][]string, 0)
	currentFriends, ok := friendships[username]

	if !ok {
		return nil
	}

	for _, friend := range currentFriends {
		mutualFriends, _ := friendships[friend]
		for _, mutFriend := range mutualFriends {
			if mutFriend == username || lo.Contains(currentFriends, mutFriend) {
				continue
			}
			suggestedFriends[mutFriend] = append(suggestedFriends[mutFriend], friend)
		}
	}

	if len(suggestedFriends) == 0 {
		return nil
	} else {
		return lo.Keys(suggestedFriends)
	}
}

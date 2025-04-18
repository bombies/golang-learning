package maps

import "errors"

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if !user.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

package arrays

import (
	"log"

	"github.com/dlclark/regexp2"
)

func isValidPassword(password string) bool {
	const REGEX string = `^(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{5,12}$`
	re := regexp2.MustCompile(REGEX, regexp2.None)

	if match, err := re.MatchString(password); err != nil {
		log.Fatalln(err)
		return false
	} else {
		return match
	}
}

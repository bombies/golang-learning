package structs

func (user User) sendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= user.MessageCharLimit {
		return message, true
	}
	return "", false
}

func TestSendMessage() {
	println(newUser("Ajani", "premium").sendMessage("Hello beautiful world", 20))
}

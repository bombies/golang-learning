package arrays

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	taggedMessages := make([]sms, 0, len(messages))
	for _, msg := range messages {
		msg.tags = tagger(msg)
		taggedMessages = append(taggedMessages, msg)
	}
	return taggedMessages
}

func tagger(msg sms) []string {
	tags := []string{}
	msgContent := strings.ToLower(msg.content)

	if strings.Contains(msgContent, "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(msgContent, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

package concurrency

import "time"

func processMessages(messages []string) []string {
	if len(messages) == 0 {
		return []string{}
	}

	msgChan := make(chan string)
	processedMsgs := make([]string, 0, len(messages))

	for _, msg := range messages {
		go func() {
			msgChan <- process(msg)
		}()
	}

	for msg := range msgChan {
		processedMsgs = append(processedMsgs, msg)

		if len(processedMsgs) == len(messages) {
			close(msgChan)
		}
	}

	return processedMsgs
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

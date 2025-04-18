package loops

func countConnections(groupSize int) int {
	res := 0
	for i := range groupSize {
		res += groupSize - (i + 1)
	}

	return res
}

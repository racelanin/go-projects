package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	novohead := l.Head

	for novohead != nil {
		if comp(novohead.Data, ref) {
			return &novohead.Data
		}

		novohead = novohead.Next
	}

	return nil
}

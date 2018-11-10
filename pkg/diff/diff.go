package diff

func DiffState(state []string, newState []string) (new []string, old []string) {

	found := false

	// Filter items only in `state`
	for i := range state {
		found = false
		for j := range newState {
			if (state[i] == newState[j]) {
				found = true
				break
			}
		}
		if (found == false) {
			old = append(old, state[i])
		}
	}

	// Filter items only in `newState`
	for j := range newState {
		found = false
		for i := range state {
			if (newState[j] == state[i]) {
				found = true
				break
			}
		}
		if (found == false) {
			new = append(new, newState[j])
		}
	}
	return new, old
}
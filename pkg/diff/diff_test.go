package diff

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type testDiffStateTable struct {
	state 		[]string
	newState 	[]string
	expectedNew []string
	expectedOld []string
}

func TestDiffState(t *testing.T) {
	tables := []testDiffStateTable {
		{[]string{}, []string{}, []string{}, []string{}},
		{
			state: []string{"apple"},
			newState: []string{"apple", "pear", "grapefruit"},
			expectedNew: []string{"pear", "grapefruit"},
			expectedOld: []string{},
		},
		{
			state: []string{"apple", "kiwi", "avocado"},
			newState: []string{},
			expectedNew: []string{},
			expectedOld: []string{"apple", "kiwi", "avocado"},
		},
		{
			state: []string{"apple", "pear"},
			newState: []string{"apple", "banana"},
			expectedNew: []string{"banana"},
			expectedOld: []string{"pear"},
		},
	}

	for _, table := range tables {
		new, old := DiffState(table.state, table.newState)
		fmt.Println(new, old)
		assert.ElementsMatch(t, new, table.expectedNew)
		assert.ElementsMatch(t, old, table.expectedOld)
	}
}

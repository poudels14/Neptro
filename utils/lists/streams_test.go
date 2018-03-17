package lists

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterStringList(t *testing.T) {
	list := []string{"hello", "world", "asdf", "kanye"}
	predicate := func(val string) bool { return val != "asdf" }
	filtered := FilterStringList(list, predicate)
	assert.Equal(t, []string{"hello", "world", "kanye"}, filtered)
}

func TestRemoveEmptyStrings(t *testing.T) {
	list := []string{"hello", "", "world", " ", "    "}
	filtered := RemoveEmptyStrings(list)
	assert.Equal(t, []string{"hello", "world"}, filtered)
}

func TestMappingStringList(t *testing.T) {
	list := []string{"hello", "world", "kanye"}
	mapper := func(val string) string { return strings.ToUpper(val) }
	mapped := MapStringList(list, mapper)
	assert.Equal(t, []string{"HELLO", "WORLD", "KANYE"}, mapped)
}

func TestFoldingStringList(t *testing.T) {
	list := []string{"hello", "mr", "west"}
	fold := func(a, b string) string { return a + b }
	folded := FoldStringList(list, "", fold)
	assert.Equal(t, "hellomrwest", folded)
}

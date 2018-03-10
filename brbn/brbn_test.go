package brbn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBrbn(t *testing.T) {
	b := New("0.0.0.0", "6666")
	assert.NotNil(t, b)
}

package brbn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBrbn(t *testing.T) {
	b := Default()
	assert.NotNil(t, b)
}

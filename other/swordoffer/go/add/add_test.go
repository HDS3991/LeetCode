package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 0, add(0, 0))
	assert.Equal(t, 8, add(3, 5))
}

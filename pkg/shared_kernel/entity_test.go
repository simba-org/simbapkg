package sharedkernel_test

import (
	"testing"

	shared "simbapkg/pkg/shared_kernel"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	id := shared.NewID()
	assert.NotNil(t, id)
}

func TestStringToID(t *testing.T) {
	t.Parallel()

	_, err := shared.StringToID("fd14c028-5f56-488a-8c29-3186fd62395c")
	assert.Nil(t, err)
}

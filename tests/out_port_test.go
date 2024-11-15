package tests

import (
	"ports/internal/services/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMultipleData(t *testing.T) {
	t.Parallel()

	outPort := ports.NewOutPort(1)

	assert.NotNil(t, outPort.GetDataChan())
}

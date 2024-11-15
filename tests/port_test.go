package tests

import (
	"ports/internal/services/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	t.Parallel()

	outPort := ports.NewPort(5, 5)
	portOut, _ := outPort.GetInPort(1)
	portIn, _ := outPort.GetOutPort(1)

	assert.NotNil(t, portOut)
	assert.NotNil(t, portIn)
}

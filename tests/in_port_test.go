package tests

import (
	"ports/internal/services/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInPort(t *testing.T) {
	portNumber := 1
	inPort := ports.NewInPort(portNumber)

	assert.NotNil(t, inPort.ResultChan())

	for i := 0; i < 40; i++ {
		assert.LessOrEqual(t, <-inPort.ResultChan(), 1)
	}

}

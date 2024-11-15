package ports

import (
	"golang.org/x/exp/rand"
)

// InPort defines the interface for an input port.
type InPort interface {
	ResultChan() <-chan int
}

// inPort implements the InPort interface.
type inPort struct {
	port
	resultChan chan int
}

// NewInPort creates a new instance of inPort and starts its work in a goroutine.
func NewInPort(portNumber int) InPort {
	inPort := &inPort{
		resultChan: make(chan int),
	}

	go func() {
		for {
			inPort.StartWork()
		}
	}()

	return inPort
}

// StartWork generates a random number (0 or 1) and sends it to the result channel.
func (p *inPort) StartWork() {
	p.resultChan <- rand.Intn(2)
}

// ResultChan returns the channel for receiving results from the inPort.
func (p *inPort) ResultChan() <-chan int {
	return p.resultChan
}

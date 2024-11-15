package ports

import (
	"fmt"
	"sync"
)

// ErrPortAbsent is an error returned when a requested port does not exist.
var ErrPortAbsent = func(number int) error {
	return fmt.Errorf("there is no port with number %d", number)
}

// Port defines the interface for managing input and output ports.
type Port interface {
	GetInPort(number int) (InPort, error)
	GetOutPort(number int) (OutPort, error)
}

// port implements the Port interface, holding maps of input and output ports.
type port struct {
	in  map[int]InPort
	out map[int]OutPort
}

// NewPort initializes a new port with the specified number of input and output ports.
func NewPort(inPortsCount, outPortsCount int) *port {
	inPorts := make(map[int]InPort)
	outPorts := make(map[int]OutPort)
	var wg sync.WaitGroup
	var mu sync.Mutex

	if inPortsCount < 1 {
		inPortsCount = 1
	}
	if outPortsCount < 1 {
		outPortsCount = 1
	}

	// Create input ports concurrently.
	for i := 1; i <= inPortsCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			inPorts[i] = NewInPort(i)
		}(i)
	}

	// Create output ports concurrently.
	for i := 1; i <= outPortsCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			outPorts[i] = NewOutPort(i)
		}(i)
	}
	wg.Wait()

	return &port{
		in:  inPorts,
		out: outPorts,
	}
}

// GetInPort retrieves an input port by its number, returning an error if not found.
func (p *port) GetInPort(number int) (InPort, error) {
	val, ok := p.in[number]
	if !ok {
		return nil, ErrPortAbsent(number)
	}

	return val, nil
}

// GetOutPort retrieves an output port by its number, returning an error if not found.
func (p *port) GetOutPort(number int) (OutPort, error) {
	val, ok := p.out[number]
	if !ok {
		return nil, ErrPortAbsent(number)
	}

	return val, nil
}

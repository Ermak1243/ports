package ports

import (
	"fmt"
	"ports/internal/models"
)

// OutPort defines the interface for an output port.
type OutPort interface {
	startListen()
	Send(number, transactionID int)
	GetDataChan() chan models.DataModel
}

// outPort implements the OutPort interface.
type outPort struct {
	port
	dataChan chan models.DataModel
}

// NewOutPort creates a new instance of outPort and starts listening for data in a goroutine.
func NewOutPort(number int) OutPort {
	outPort := &outPort{
		dataChan: make(chan models.DataModel),
	}

	// Start a goroutine to listen for incoming data.
	go func() {
		for {
			outPort.startListen()
		}
	}()

	return outPort
}

// startListen waits for incoming data on the data channel and processes it.
func (p *outPort) startListen() {
	data := <-p.dataChan
	fmt.Printf("OUT port %d: transaction ID %d\n", data.Port, data.TransactionID)
}

func (p *outPort) GetDataChan() chan models.DataModel {
	return p.dataChan
}

// Send sends a DataModel containing the port number and transaction ID to the output port.
func (p *outPort) Send(number, transactionID int) {
	p.dataChan <- models.DataModel{
		Port:          number,
		TransactionID: transactionID,
	}
}

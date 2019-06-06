package structures

import (
	"fmt"
	"go-dp/observer-dp/interfaces"
	"strings"
)

// StockGrabber structure
type StockGrabber struct {
	SBI       int
	HDFC      int
	IDFC      int
	Observers []interfaces.Observer
}

// UpdateSBI updates SBI value
func (s *StockGrabber) UpdateSBI(SBI int) {
	s.SBI = SBI
	fmt.Println("SBI Updated..")
	s.NotifyUsers()
}

// UpdateHDFC updates HDFC value
func (s *StockGrabber) UpdateHDFC(HDFC int) {
	s.HDFC = HDFC
	fmt.Println("HDFC Updated..")
	s.NotifyUsers()
}

// UpdateIDFC updates IDFC value
func (s *StockGrabber) UpdateIDFC(IDFC int) {
	s.IDFC = IDFC
	fmt.Println("IDFC Updated..")
	s.NotifyUsers()
}

// Update implementation from Subject Interface
func (s *StockGrabber) Update(SBI, HDFC, IDFC int) {
	s.SBI = SBI
	s.HDFC = HDFC
	s.IDFC = IDFC
	fmt.Println("All Stocks Updated..")
	s.NotifyUsers()
}

// AddObserver adds the Observer object to the StockGrabber
func (s *StockGrabber) AddObserver(o interfaces.Observer) {
	s.Observers = append(s.Observers, o)
}

// DeleteObserver deletes the Observer object to the StockGrabber
func (s *StockGrabber) DeleteObserver(o interfaces.Observer) {
	index := -1
	for i, existingObserver := range s.Observers {
		if existingObserver == o {
			index = i
		}
	}

	if index != -1 {
		s.Observers = append(s.Observers[:index], s.Observers[index+1:]...)
	}
}

// NotifyUsers calls the Notify() method on Observers
func (s *StockGrabber) NotifyUsers() {
	for _, o := range s.Observers {
		o.Notify(s.SBI, s.HDFC, s.IDFC)
	}
}

func (s StockGrabber) String() string {
	stringRep := []string{"StockGrabber", "Observers"}
	for _, o := range s.Observers {
		stringRep = append(stringRep, o.String())
	}
	return strings.Join(stringRep, "\n")
}

// NewStockGrabber for creating the StockGrabber
func NewStockGrabber(SBI, HDFC, IDFC int) StockGrabber {
	return StockGrabber{SBI: SBI, HDFC: HDFC, IDFC: IDFC}
}

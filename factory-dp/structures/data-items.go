package structures

import (
	"fmt"
)

// Memory Structure
type Memory struct {
	Available int
	Used      int
	Free      int
}

// GetDetails implementation for Memory
func (i Memory) GetDetails() string {
	return fmt.Sprintf("Memory Details\nAvailable %v \nUsed %v \nFree %v", i.Available, i.Used, i.Free)
}

// CPU Structure
type CPU struct {
	Core1 int
	Core2 int
	Core3 int
}

// GetDetails implementation for CPU
func (i CPU) GetDetails() string {
	return fmt.Sprintf("CPU Details\nCore1 %v\nCore2 %v\nCore3 %v", i.Core1, i.Core2, i.Core3)
}

// Monitor Structure
type Monitor struct {
	Resolution  string
	RefreshRate int
}

// GetDetails implementation for Monitor
func (i Monitor) GetDetails() string {
	return fmt.Sprintf("Monitor Details\nResolution %v\nRefresh Rate %v", i.Resolution, i.RefreshRate)
}

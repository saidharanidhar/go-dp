package interfaces

import (
	"fmt"
)

// Observer Interface
type Observer interface {
	Notify(SBI, HDFC, IDFC int)
	fmt.Stringer
}

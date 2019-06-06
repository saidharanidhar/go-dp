package interfaces

import (
	"fmt"
)

// Subject Interface
type Subject interface {
	Update(SBI, HDFC, IDFC int)
	fmt.Stringer
}

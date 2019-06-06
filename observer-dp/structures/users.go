package structures

import (
	"fmt"
	"strconv"
)

// User Structure
type User struct {
	id int
}

// Notify implementation from Observer interface
func (u *User) Notify(SBI, HDFC, IDFC int) {
	fmt.Println("User: ", u.id, "SBI: ", SBI, "HDFC: ", HDFC, "IDFC: ", IDFC)
}

func (u User) String() string {
	return "User: " + strconv.Itoa(u.id)
}

// Closure for maintaing static count
func newUser() func() *User {
	staticID := -1
	return func() *User {
		staticID++
		return &User{staticID}
	}
}

// NewUser creates User and auto assigns id
var NewUser = newUser()

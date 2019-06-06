package main

import (
	"fmt"
	"go-dp/observer-dp/structures"
)

func main() {
	stockgrabber := structures.NewStockGrabber(0, 0, 0)

	newUser := structures.NewUser()
	stockgrabber.AddObserver(newUser)

	stockgrabber.Update(1, 2, 3)

	newUser2 := structures.NewUser()
	stockgrabber.AddObserver(newUser2)

	stockgrabber.Update(3, 4, 5)

	stockgrabber.DeleteObserver(newUser)

	stockgrabber.Update(5, 6, 7)
	stockgrabber.UpdateHDFC(9)

	newUser3 := structures.NewUser()
	stockgrabber.AddObserver(newUser3)

	stockgrabber.UpdateSBI(12)

	fmt.Println(stockgrabber)

}

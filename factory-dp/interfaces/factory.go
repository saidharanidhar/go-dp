package interfaces

// Factory interface
type Factory interface {
	GetItem(string) (Item, error)
	Describe()
}

package structures

import (
	"fmt"
	"go-dp/factory-dp/interfaces"
)

// DataFactory Structure
type DataFactory struct {
	ItemsMap map[string]interfaces.Item
}

// GetItem returns the item if present
func (factory DataFactory) GetItem(id string) (interfaces.Item, error) {
	item, ok := factory.ItemsMap[id]
	if !ok {
		return item, fmt.Errorf("Item with id %v not found", id)
	}
	return item, nil
}

// AddToFactory adds items to given string
func (factory *DataFactory) AddToFactory(id string, item interfaces.Item) {
	factory.ItemsMap[id] = item
}

// Describe the Factory
func (factory DataFactory) Describe() {
	fmt.Println("Available Items in the Factory:")
	for i := range factory.ItemsMap {
		fmt.Println(i)
	}
}

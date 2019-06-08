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

// addToFactory adds items to given string
func (factory *DataFactory) addToFactory(id string, item interfaces.Item) {
	factory.ItemsMap[id] = item
}

// Describe the Factory
func (factory DataFactory) Describe() {
	fmt.Println("Available Items in the Factory:")
	for i := range factory.ItemsMap {
		fmt.Println(i)
	}
}

// NewFactory Creates a New Factory structure
func NewFactory() interfaces.Factory {
	factory := DataFactory{
		ItemsMap: make(map[string]interfaces.Item),
	}

	factory.addToFactory(
		"cpu",
		interfaces.Item(
			CPU{
				Core1: 12,
				Core2: 14,
				Core3: 19,
			}))

	factory.addToFactory(
		"memory",
		interfaces.Item(
			Memory{
				Available: 12,
				Used:      14,
				Free:      16,
			}))

	factory.addToFactory(
		"monitor",
		interfaces.Item(
			Monitor{
				Resolution:  "1900x1080",
				RefreshRate: 195,
			}))

	return interfaces.Factory(factory)

}

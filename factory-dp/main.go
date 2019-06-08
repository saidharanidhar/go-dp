package main

import (
	"bufio"
	"fmt"
	"go-dp/factory-dp/interfaces"
	"go-dp/factory-dp/structures"
	"os"
	"strings"
)

func main() {
	var factory interfaces.Factory = &structures.DataFactory{
		ItemsMap: make(map[string]interfaces.Item),
	}

	var cpu interfaces.Item = structures.CPU{
		Core1: 12,
		Core2: 14,
		Core3: 19,
	}

	var memory interfaces.Item = structures.Memory{
		Available: 12,
		Used:      14,
		Free:      16,
	}
	var monitor interfaces.Item = structures.Monitor{
		Resolution:  "1900x1080",
		RefreshRate: 195,
	}

	factory.AddToFactory("cpu", cpu)
	factory.AddToFactory("memory", memory)
	factory.AddToFactory("monitor", monitor)

	factory.Describe()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select from above > ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	item, err := factory.GetItem(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item.GetDetails())
}

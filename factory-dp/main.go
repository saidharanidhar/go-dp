package main

import (
	"bufio"
	"fmt"
	"go-dp/factory-dp/structures"
	"os"
	"strings"
)

func main() {
	factory := structures.NewFactory()
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

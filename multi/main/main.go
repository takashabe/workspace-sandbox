package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/multi/multishop"
)

func main() {
	s := &multishop.Shop{
		Name: "multi",
	}
	fmt.Printf("%#v", s)
}

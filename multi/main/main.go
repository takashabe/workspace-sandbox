package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/multi/shop"
)

func main() {
	s := &shop.Shop{
		Name: "multi",
	}
	fmt.Printf("%#v", s)
}

package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/workspace/shop"
)

func main() {
	s := &shop.Shop{
		Name: "workspace",
	}
	fmt.Printf("%#v", s)
}

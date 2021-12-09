package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/workspace/workspaceshop"
)

func main() {
	s := &shop.Shop{
		Name: "multi",
	}
	fmt.Printf("%#v", s)
}

package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/workspace/workspaceshop"
)

func main() {
	s := &workspaceshop.Shop{
		Name: "multi",
	}
	fmt.Printf("%#v", s)
}

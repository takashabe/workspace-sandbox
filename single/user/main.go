package main

import (
	"fmt"

	singleshop "github.com/takashabe/workspace-sandbox/single/shop"
)

func main() {
	s := &singleshop.Shop{
		Name: "single",
	}
	fmt.Println(s)
}

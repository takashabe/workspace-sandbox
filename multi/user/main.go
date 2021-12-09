package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/workspace/shop"
)

func main() {
	u := &User{
		Favolite: &shop.Shop{
			Name:     "new",
			Address:  "tokyo",
			Category: "book",
		},
	}
	u.fav()
}

type User struct {
	Email    string
	Name     string
	Favolite *shop.Shop
}

func (u *User) fav() {
	fmt.Printf("favolite category: %s", u.Favolite.Category)
}
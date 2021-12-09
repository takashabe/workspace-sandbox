package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/multi/mshop"
)

func main() {
	u := &User{
		Favolite: &mshop.Shop{
			Name:     "new",
			Address:  "tokyo",
			Category: "aaa",
		},
	}
	u.fav()
}

type User struct {
	Email    string
	Name     string
	Favolite *mshop.Shop
}

func (u *User) fav() {
	fmt.Printf("favolite category: %v", u.Favolite.Category)
}

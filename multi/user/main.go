package main

import (
	"fmt"

	"github.com/takashabe/workspace-sandbox/multi/mshop"
)

func main() {
	u := &User{
		Favolite: &mshop.Shop{
			Name:    "new",
			Address: "tokyo",
			Cate:    1,
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
	fmt.Printf("favolite category: %v", u.Favolite.Cate)
}

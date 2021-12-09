package main

import (
	"github.com/takashabe/workspace-sandbox/workspace/shop"
)

func main() {
	u := &User{
		Favolite: &shop.Shop{
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
	Favolite *shop.Shop
}

func (u *User) fav() {
	// fmt.Printf("favolite category: %s", u.Favolite.Category)
}

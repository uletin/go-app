package main

import (
	"fmt"

	"github.com/uletin/go-app/helpers"
)

func main() {
	c := helpers.NewCustomer(helpers.WithFullname("Etika"))
	fmt.Println(c)
}

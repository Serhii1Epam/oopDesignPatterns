package main

import (
	"fmt"

	"github.com/Serhii1Epam/oopDesignPatterns/internal/pkg/factory"
)

func main() {
	hd, _ := factory.GetMoto("hd")
	bmw, _ := factory.GetMoto("bmw")

	fmt.Println(hd)
	fmt.Println(bmw)
}

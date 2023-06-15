package main

import (
	"fmt"

	"github.com/Serhii1Epam/oopDesignPatterns/internal/pkg/builder"
)

func main() {
	computerBuilder := builder.NewComputerBuilder()
	computer := computerBuilder.CPU("core_i7").RAM(16).MB("GIGABYTE").Build()
	fmt.Println(computer)

	officeComputerBuilder := builder.NewOfficeCumputerBuilder()
	officeComputer := officeComputerBuilder.Build()
	fmt.Println(officeComputer)
}

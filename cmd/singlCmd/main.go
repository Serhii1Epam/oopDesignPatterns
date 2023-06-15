package main

import (
	"fmt"

	"github.com/Serhii1Epam/oopDesignPatterns/internal/pkg/singleton"
)

func main() {
	var snglt singleton.SingletonUniqueId
	s := snglt.GetUniqueId()
	buildSmth1(s.Next())
	buildSmth1(s.Next())
	s1 := snglt.GetUniqueId()
	buildSmth2(s1.Next())
	buildSmth2(s1.Next())
}

func buildSmth1(id int) {
	fmt.Printf("buildSmth1, id:%d\n", id)

}

func buildSmth2(id int) {
	fmt.Printf("buildSmth2, id:%d\n", id)

}

package singleton

import (
	"fmt"
)

type SingletonUniqueId struct {
	UniqueId *UniqueId
}

func NewUniqueId() *SingletonUniqueId {
	return &SingletonUniqueId{}
}

func (s *SingletonUniqueId) GetUniqueId() *UniqueId {
	if s.UniqueId == nil {
		fmt.Println("UniqueId not initalized. Create it.")
		s.UniqueId = New()
	}
	return s.UniqueId
}

type UniqueId struct {
	unique int
}

func New() *UniqueId {
	return &UniqueId{0}
}

func (u *UniqueId) Next() int {
	u.unique++
	return u.unique
}

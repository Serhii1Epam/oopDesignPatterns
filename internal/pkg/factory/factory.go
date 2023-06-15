package factory

import "fmt"

type IMoto interface {
	Model() string
	Using() string
	SetModel(val string)
	SetUsing(val string)
}

type moto struct {
	model, using string
}

func (m *moto) Model() string {
	return m.model
}
func (m *moto) Using() string {
	return m.using
}

func (m *moto) SetModel(val string) {
	m.model = val
}
func (m *moto) SetUsing(val string) {
	m.using = val
}

type CruiserMoto struct {
	moto
}

func NewCruiserMoto() IMoto {
	return &CruiserMoto{
		moto: moto{
			model: "Harley-Davidson RoadGlide",
			using: "cruiser",
		},
	}
}

type RaceMoto struct {
	moto
}

func NewRaceMoto() IMoto {
	return &RaceMoto{moto{"BMW 1000RR", "racer"}}
}

func GetMoto(moto string) (IMoto, error) {
	if moto == "hd" {
		return NewCruiserMoto(), nil
	}
	if moto == "bmw" {
		return NewRaceMoto(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

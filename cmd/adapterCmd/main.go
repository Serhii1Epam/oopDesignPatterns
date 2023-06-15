package main

import (
	"github.com/Serhii1Epam/oopDesignPatterns/internal/pkg/adapter"
)

func main() {
	//first/implementation
	json := &adapter.JsonData{}
	client := &adapter.Client{}
	client.Send(json)

	//Extended implementation
	xml := &adapter.XmlData{}
	xmlAd := &adapter.XmlAdapter{XmlData: xml}
	client.Send(xmlAd)
}

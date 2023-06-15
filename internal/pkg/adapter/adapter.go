package adapter

import "fmt"

// Target
type Sender interface {
	SendJson()
}

// Concrete prototype implementation
type JsonData struct {
}

func (j *JsonData) SendJson() {
	fmt.Println("Send JSON data to somewhere.")
}

// Client
type Client struct {
}

func (c *Client) Send(s Sender) {
	s.SendJson()
}

// Adaptee
type XmlData struct {
}

func (x *XmlData) SendXml() {
	fmt.Println("Send XML data to somewhere.")
}

// Adapter
type XmlAdapter struct {
	XmlData *XmlData
}

func (xa *XmlAdapter) SendJson() {
	xa.XmlData.SendXml()
}

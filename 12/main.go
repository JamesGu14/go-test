package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected: ", pc.name)
}

func main() {
	var a USB = PhoneConnector{"My Phone Connector"}
	a.Connect()
	Disconnect(a)
	var b USB
	Disconnect(b)
	fmt.Println("===============")
	Disconnect2(a)
	Disconnect2(b)
}

func Disconnect(usb USB) {

	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("Disconnected. ", pc.Name())
		return
	}

	fmt.Println("Unknown device.")
}

func Disconnect2(usb interface{}) {

	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected2: " + v.Name())
	default:
		fmt.Println("Unknown Device2.")
	}
}

package main

import (
	"gate"
)

func main() {
	loginServerObject := gate.SingleGateServer()
	loginServerObject.Start()
}

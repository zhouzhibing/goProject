package net

import (
	"sync"
	"gameTools/game/tools/net/sockets/websocket"
	"gameTools/game/tools/config"

	"strconv"
)

type socketControl struct {
	webSocket * websocket.WebSocketServer
}

var socketControlObject * socketControl
var lock sync.Mutex

func SingleSocketControl() * socketControl{

	if socketControlObject != nil {
		return socketControlObject
	}

	lock.Lock()
	defer lock.Unlock()

	if socketControlObject != nil {
		return socketControlObject
	}

	if socketControlObject == nil {
		socketControlObject = new(socketControl)
		socketControlObject.init()
		return socketControlObject
	}

	return socketControlObject
}

func (this *socketControl) init() {
	port := config.GetSingleConfigValue("server.port")
	portInt , err := strconv.Atoi(port)
	if err != nil{
		panic(err)
	}

	this.webSocket = websocket.NewWebSocketServer(portInt, new(SocketHandle))
}


func (this *socketControl) Start() {
	this.webSocket.Start()
}
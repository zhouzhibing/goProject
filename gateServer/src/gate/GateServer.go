package gate

import (
	"gameTools/game/tools"
	"gameTools/game/tools/config"
	"gate/control"
	"gate/db/redis"
	"gate/net"
	"os"
	"sync"
	"fmt"
	"runtime"
)

type gateServer struct{}

var gateServerObject *gateServer
var lock * sync.Mutex = tools.NewMutex()

func SingleGateServer() *gateServer {

	if gateServerObject != nil {
		return gateServerObject
	}

	lock.Lock()
	defer lock.Unlock()

	if gateServerObject != nil {
		return gateServerObject
	}

	if gateServerObject == nil {
		gateServerObject = new(gateServer)
		return gateServerObject
	}

	return gateServerObject
}

func (this *gateServer) Start() {
	this.initConf()
	this.initRedis()
	this.initControl()
	this.initNet()
	fmt.Println("RunGateServer Start!")
	this.initCmd()
}
func (server *gateServer) initCmd() {

	tools.DoCmd(func(text string) {
		switch text {
			case"gocount":
				fmt.Println("go:" , runtime.NumGoroutine())
			case "stop":
			os.Exit(-1)
		}
		if text == "stop" {
			os.Exit(-1)
		}
	})
}

func (this *gateServer) initControl() {
	gateControl := control.SingleGateControl()
	gateControl.Start()
}

func (this *gateServer) initRedis() {
	redis.RedisControlInit()
	redis.RegisterGateServer()
	redis.InitSceneList()
}

func (this *gateServer) initNet() {
	socketControl := net.SingleSocketControl()
	go socketControl.Start()
}

func (this *gateServer) initConf() {
	config.SingleConfig("gateServer/conf/gateServer.conf")
}

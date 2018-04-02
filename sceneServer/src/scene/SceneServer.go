package scene

import (
	"sync"
	"gameTools/game/tools/config"
	"scene/game/net"
	"scene/game/db/redis"
	"strconv"
	"scene/game/msg/api"
	"scene/game/control"
	"scene/game/db/mysql"
	"os"
	"gameTools/game/tools"
	"runtime"
	"fmt"
)

type sceneServer struct {}

var sceneServerObject * sceneServer
var sceneLock sync.Mutex

func SingleSceneServer() * sceneServer{

	if sceneServerObject != nil {
		return sceneServerObject
	}

	sceneLock.Lock()
	defer sceneLock.Unlock()

	if sceneServerObject != nil {
		return sceneServerObject
	}

	if sceneServerObject == nil {
		sceneServerObject = new(sceneServer)
		return sceneServerObject
	}

	return sceneServerObject
}

func (this * sceneServer) Start(){
	initConf()
	initMysql()
	initRedis()
	initWorld()
	initApi()
	initNet()
	initCmd()
}

func initCmd(){
	tools.DoCmd(func(text string) {
		switch text {
		case "stop":
			os.Exit(-1);
		case "qps":
			fmt.Println("qps : " , control.SingleWorldManager().GetQpsCount())
		case "tps":
			fmt.Println("tps : " , control.SingleWorldManager().GetTpsCount())
		case "gocount":
			fmt.Println("go:" , runtime.NumGoroutine())
		case "online":
			fmt.Println("online:" , control.SingleWorldManager().UserCount())
		}
		if text == "stop" {
			os.Exit(-1)
		}
	})
}

func initMysql() {
	mysql.MysqlControlInit()
}

func initApi() {
	api.SingleApi().InitApi()
}

func initWorld() {
	worldManager := control.SingleWorldManager()
	worldManager.Start()
}

func initRedis() {
	redis.RedisControlInit()
	redis.RegisterSceneServer()
}

func initNet() {
	port := config.GetSingleConfigValue("server.port")
	portInt, err := strconv.Atoi(port)
	if err != nil{
		panic(err)
	}

	socketControl := net.NewNetServer(portInt)
	go socketControl.Start()
}

func initConf() {
	config.SingleConfig("sceneServer/conf/sceneServer.conf")
}


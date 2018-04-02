package websocket

import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"strconv"
	"time"
	"runtime"
	"gameTools/game/tools"
)


//超时时间，单位秒
const TIME_OUT  = 10
const ROOT_PATH = "WebSocket"

type WebSocketServer struct {
	port int

	rootPath string

	server *http.Server

	handle IWebSocketHandle													//通信逻辑处理对象
}

func NewWebSocketServer(port int, handle IWebSocketHandle) *WebSocketServer {
	runtime.GOMAXPROCS(runtime.NumCPU())

	this := new(WebSocketServer)
	this.handle = handle

	this.rootPath = "/" + ROOT_PATH
	this.port = port
	this.init()

	return this
}

func (this *WebSocketServer) init(){

	portString := ":" + strconv.Itoa(this.port)

	this.server = new(http.Server)
	this.server.Addr = portString
	this.server.ReadTimeout =  TIME_OUT * time.Second
	this.server.WriteTimeout =  TIME_OUT * time.Second
	this.server.SetKeepAlivesEnabled(true)
}

func (this *WebSocketServer) Start(){

	fmt.Println("HttpServer.Start by : " , this.rootPath  ,":", this.port )

	defer tools.Exception(func(errObject interface{}) {
		fmt.Println("Exception HttpServer Start " , errObject)
	})

	http.Handle("/", websocket.Handler(this.handler))

	err := this.server.ListenAndServe()

	if err != nil{
		panic(err)
	}
}

func (this * WebSocketServer) handler(conn *websocket.Conn){
	var err error

	//fmt.Println("WebSocketHandle.Handle connection! " , this)
	connection := NewConnection(conn)
	this.handle.Connection(connection)

	for {

		var reply string

		if err = websocket.Message.Receive(conn, &reply); err != nil {			//有错误，断开了连接
			//fmt.Println("Can't receive")
			//panic(err)
			connection.Close()
			this.handle.Close(connection)
			break
		}

		this.handle.Handle(connection , reply)
	//	fmt.Println("Received back from client: " + reply)

		//msg := "Received:  " + reply
		//connection.Send(msg)
		//fmt.Println("Sending to client: " + msg)

		//if err = websocket.Message.Send(conn, msg); err != nil {
		//	fmt.Println("Can't send")
		//	break
		//}
	}
}

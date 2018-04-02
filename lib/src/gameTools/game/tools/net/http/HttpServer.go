package http

import (
	"net/http"
	"fmt"
	"strconv"
	"time"
	"gameTools/game/tools/collect"
	"runtime"
	"reflect"
	"gameTools/game/tools"
)

//超时时间，单位秒
const TIME_OUT  = 10

type HttpServer struct {
	rootPath string
	port int
	server *http.Server
}

func NewHttpServer(rootPath string , port int ) * HttpServer{
	runtime.GOMAXPROCS(runtime.NumCPU())

	this := new(HttpServer)
	this.rootPath = "/" + rootPath
	this.port = port
	this.init()

	return this
}


func (this * HttpServer) init(){

	portString := ":" + strconv.Itoa(this.port)

	this.server = new(http.Server)
	this.server.Addr = portString
	this.server.ReadTimeout =  TIME_OUT * time.Second
	this.server.WriteTimeout =  TIME_OUT * time.Second
	this.server.SetKeepAlivesEnabled(true)
}

func (this * HttpServer) Start(){

	fmt.Println("HttpServer.Start by : " , this.rootPath  ,":", this.port )

	defer tools.Exception(func(errObject interface{}) {
		fmt.Println("Exception HttpServer Start " , errObject)
	})

	err := this.server.ListenAndServe()

	if err != nil{
		panic(err)
	}
}

func (this * HttpServer) RegisterHandle(httpHande IHttpHandle){

	patten := reflect.ValueOf(httpHande).Elem().Type().Name()
	path := this.rootPath + "/" + patten

	http.DefaultServeMux.HandleFunc(path, httpHande.Handle)

	fmt.Println("RegisterHandle : " , path)
}

func (this * HttpServer) RegisterHandles(handleList * collect.ArrayList){
	handleList.Pool(func(index interface{}, value interface{}) {
		this.RegisterHandle(value.(IHttpHandle))
	})
}

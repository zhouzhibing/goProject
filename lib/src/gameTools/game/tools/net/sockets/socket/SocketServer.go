package socket

import (
	"net"
	"fmt"
	"strconv"
	"runtime"
	"bytes"
)

type SocketServer struct {

	port int														//port端口

	netListen net.Listener											//net通信处理对象

	deencode IDeEncode												//所需要的编码解码对象

	handle ISocketHandle											//通信逻辑处理对象
}
/**
*	port 监听端口
*   deencode 编码解码器
*   handle 处理器
*/
func NewSocketServer(port int , deencode IDeEncode , handle ISocketHandle) * SocketServer{

	runtime.GOMAXPROCS(runtime.NumCPU())

	this := new(SocketServer)

	this.port = port
	this.deencode = deencode
	this.handle = handle

	return this
}

func (this * SocketServer) Start(){
	this.accept()
}

func (this * SocketServer) accept(){

	if this.netListen != nil{
		return
	}

	addr := "localhost:" + strconv.Itoa(this.port);
	netListen, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			fmt.Println( " netListen.Accept conn :" , conn , " err : ",err)
			continue
		}
		//fmt.Println(conn.RemoteAddr().String() + " tcp connect success")

		connection := NewConnection(conn , this.deencode)

		this.handle.Connection(connection)

		go this.handleConnection(connection)

		fmt.Println("SocketServer Go Routine Count : " , runtime.NumGoroutine())
	}
}


func (this * SocketServer) handleConnection(conn* Connection){

	//fmt.Println(conn.RemoteAddr())

	buffer := new(bytes.Buffer)

	readerBuffer := make([]byte , 5024)

	for {

		readerLegnth, err := conn.Read(readerBuffer)

		if err != nil {
			//fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			conn.Close()
			this.handle.Close(conn)
			return
		}
		//fmt.Println("reader buffer size : " , readerLegnth)

		buffer.Write(readerBuffer[:readerLegnth])

		this.deencode.Decode(conn ,buffer, this.handle.Handle)

	}
}

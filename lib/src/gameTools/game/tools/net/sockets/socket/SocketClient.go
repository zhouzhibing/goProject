package socket

import (
	"net"
	"fmt"
	"os"
	"uuid-master"
	"runtime"
	"bytes"
)

type SocketClient struct {

	id string

	ipPort string

	connectionObject net.Conn										//连接对象

	deencode IDeEncode

	handleFunc func(conn * Connection , msgObject interface{})		//逻辑处理
}

func NewSocketClient(ipPort string ,deencode IDeEncode, handleFunc func(conn * Connection , msgObject interface{})) *SocketClient {
	this := new(SocketClient)

	this.ipPort = ipPort
	this.deencode = deencode
	this.handleFunc = handleFunc
	this.id = uuid.Rand().Hex()

	this.connection(ipPort , deencode , handleFunc)

	return this
}

func (this * SocketClient) ReConnection() {

	this.connection(this.ipPort , this.deencode , this.handleFunc)
}

func (this * SocketClient) connection(ip_port string ,deencode IDeEncode, handleFunc func(conn * Connection , msgObject interface{})){

	//if(this.connectionObject != nil){
	//	fmt.Println("Already connection !!")
	//	return
	//}

	if(this.connectionObject != nil){
		this.connectionObject.Close();
	}

	server := ip_port
	this.deencode = deencode
	this.handleFunc = handleFunc

	/*
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	*/
	conn, err := net.Dial("tcp", server)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	this.connectionObject = conn;

	connection := NewConnection(conn , this.deencode)

	go this.handleConnection(connection)

	fmt.Println("SocketClient Go Routine Count : " , runtime.NumGoroutine())
}

func (this * SocketClient) handleConnection(conn* Connection){

	//fmt.Println(conn.RemoteAddr())

	buffer := new(bytes.Buffer)

	readerBuffer := make([]byte , 5024)

	for {
		readerLegnth, err := conn.Read(readerBuffer)

		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		//fmt.Println("reader buffer size : " , readerLegnth)

		buffer.Write(readerBuffer[:readerLegnth])

		this.deencode.Decode(conn ,buffer, this.handleFunc)

	}
}


func (this * SocketClient) SendString(msgString string) error{

	byteArray := this.deencode.Encode(msgString)

	_ , err := this.connectionObject.Write(byteArray)

	return err
}

func (this * SocketClient) Send(byteArray [] byte){
	_ , err := this.connectionObject.Write(byteArray)

	if err != nil{
		panic(err)
	}
}

func (this *SocketClient) IsClose() {

}

func (this *SocketClient) Close() {
	this.connectionObject.Close()
}

func (this *SocketClient) GetId() string{
	return this.id
}


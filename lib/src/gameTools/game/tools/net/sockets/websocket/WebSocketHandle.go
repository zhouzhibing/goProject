package websocket

import "fmt"

type WebSocketHandle struct{

}

func (this * WebSocketHandle) Connection(connection * Connection){
	fmt.Println("new connection")
	connection.SetAttachment("asdf")
}

func (this * WebSocketHandle) Handle(connection * Connection , msgObject interface{})	{
	fmt.Println("handle msgObject = " , msgObject , " attackment : " , connection.GetAttachment() )
}

func (this * WebSocketHandle) Close(connection * Connection ){
	fmt.Println("close " , " attackment : " , connection.GetAttachment())
}
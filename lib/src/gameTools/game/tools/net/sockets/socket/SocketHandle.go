package socket

import	"fmt"

type SocketHandle struct{

}

func (this * SocketHandle) Connection(connection * Connection){
	fmt.Println("new connection")
	connection.SetAttachment("asdf")
}

func (this * SocketHandle) Handle(connection * Connection , msgObject interface{})	{
	fmt.Println("handle msgObject = " , msgObject , " attackment : " , connection.GetAttachment() )
}

func (this * SocketHandle) Close(connection * Connection ){
	fmt.Println("close " , " attackment : " , connection.GetAttachment())
}
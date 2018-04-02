package net

import (
	 "gameTools/game/tools/net/sockets/websocket"
	"fmt"
	"gate/control"
	"gameTools/game/tools/json"
	"gameTools/game/tools"
)

type SocketHandle struct {

}

func (this *SocketHandle) Connection(connection * websocket.Connection){		//新的连接
	defer tools.Exception(nil)

	fmt.Println("SocketHandle.Connection")
}

func (this *SocketHandle) Handle (connection * websocket.Connection , msgObject interface{}){			//处理连接逻辑

	defer tools.Exception(nil)

	fmt.Println("SocketHandle.Handle msgObject : " , msgObject)

	msgString := msgObject.(string)

	o := json.ParseJSONObject(msgString)

	attachment := connection.GetAttachment()

	if attachment == nil{
		attachment = this.enterGame(connection , o)
	}

	playControl := attachment.(* control.PlayControl)

	playControl.DoAction(o)

}

func (this * SocketHandle) enterGame(connection * websocket.Connection , o *json.JSONObject) * control.PlayControl{

	uId := o.GetString("uId")
	id := o.GetInt64("id")

	playControl := control.NewPlayControl(uId ,id ,connection)

	control.SingleGateControl().AddPlayControl(playControl)

	return playControl
}


func (this *SocketHandle) Close(connection * websocket.Connection){				//连接关闭
	defer tools.Exception(nil)

	//fmt.Println("SocketHandle.Close")

	connectionObject := connection.GetAttachment()

	if connectionObject != nil {

		playControl := connectionObject.(* control.PlayControl)

		playControl.Close()

		control.SingleGateControl().RemovePlayControl(playControl)

		connection.SetAttachment(nil)
	}
}
package control

import (
	"gameTools/game/tools/net/sockets/websocket"
	"gameTools/game/tools/json"
	"gameTools/game/tools/net/sockets/socket"
	"gate/db/redis"
	"fmt"
	"gameTools/game/tools/protocol"
)

type PlayControl struct {
	uId string				//用户登录ID
	id int64				//库里的唯一
	clientConnection * websocket.Connection
	sceneId int			//当前这个角色所在的场景ID
	sceneSocket * socket.SocketClient
}

func NewPlayControl(uId string , id int64 , connection * websocket.Connection) *PlayControl{
	this := new(PlayControl)
	this.uId = uId
	this.id = id
	this.clientConnection = connection
	this.init()

	return this
}

func (this * PlayControl) init(){
	this.clientConnection.SetAttachment(this)
}

func (this * PlayControl) NewConnection(){

}

func (this * PlayControl) GetUid() string {
	return this.uId
}

func (this * PlayControl) GetId() int64 {
	return this.id
}

func (this * PlayControl) Send(msgObject interface{}) {
	this.clientConnection.Send(msgObject)
}

func (this *PlayControl ) createSceneScoket(sceneId int){
	if this.sceneId == sceneId{
		return
	}
	if sceneId > 0 {
		this.sceneId = sceneId
	}

	if this.sceneSocket != nil{
		this.sceneSocket = nil
	}

	if this.sceneSocket == nil{
		this.sceneSocket = SingleGateControl().GetSceneScoket(sceneId)
	}

	sceneJson := redis.GetSceneBySceneId(sceneId)

	if sceneJson == nil{
		return
	}

	address := sceneJson.GetString("ip") + ":" + sceneJson.GetString("port")

	this.sceneSocket = socket.NewSocketClient(address , socket.NewStringDeEncode() , func(conn *socket.Connection, msgObject interface{}) {
		this.Send(msgObject)
	})
}

//切换场景socket
func (this * PlayControl) switchSceneSocket(toSceneId int){

	if this.sceneId != 0 {

		if this.sceneId != toSceneId{

			if this.sceneSocket != nil{

				exitScene := json.NewJSONObject()
				exitScene.Put("uId" , this.GetUid())
				exitScene.Put("toSceneId" , toSceneId)
				exitScene.Put(protocol.MSG_HANDLE_ID , this.GetId())
				exitScene.Put("protocolNo" , protocol.MSG_S2S_EXIT_SCENE)
				exitScene.Put("type" , "switchScene")

				this.sceneSocket.SendString(exitScene.ToJSONString())
			}
		}
	}

	this.sceneId = toSceneId

	this.sceneSocket = SingleGateControl().GetSceneScoket(this.sceneId)
}


func (this *PlayControl) DoAction(o *json.JSONObject) {

	protocolNo := o.GetInt("protocolNo")

	if protocolNo == protocol.MSG_C2S_ENTER_SCENE{

		toSceneId := o.GetInt("toSceneId")
		this.switchSceneSocket(toSceneId)
	}

	o.Put(protocol.MSG_HANDLE_ID , this.GetId())
	msgString := o.ToJSONString()

	err := this.sceneSocket.SendString(msgString)

	if err != nil{				//如果发送失败，则进行重连一次。
		this.sceneSocket.ReConnection()

		fmt.Println("SceneSocket Err : " , err)
		this.sceneSocket.SendString(msgString)
	}

	//fmt.Println("Send Scene Server Connection : " , msgString)

}

func (this *PlayControl) Close() {

	fmt.Println("id " , this.GetId() , " uId "  , this.GetUid())

	playClose:= json.NewJSONObject()
	playClose.Put("uId" , this.GetUid())
	playClose.Put("toSceneId" , 0)
	playClose.Put(protocol.MSG_HANDLE_ID , this.GetId())
	playClose.Put("protocolNo" , protocol.MSG_S2S_EXIT_SCENE)
	playClose.Put("type" , "close")

	this.sceneSocket.SendString(playClose.ToJSONString())
}
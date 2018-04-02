package control

import (
	"sync"
	"gameTools/game/tools/collect"
	"gate/db/redis"
	"gameTools/game/tools/net/sockets/socket"
	"gameTools/game/tools/json"
	"gameTools/game/tools"
	"fmt"
	"gameTools/game/tools/protocol"
)

type gateControl struct {
	playControlMap* collect.Map
	sceneSocketMap* collect.Map
}

var gateControlObject * gateControl
var lock sync.Mutex

func SingleGateControl() * gateControl{

	if gateControlObject != nil {
		return gateControlObject
	}

	lock.Lock()
	defer lock.Unlock()

	if gateControlObject != nil {
		return gateControlObject
	}

	if gateControlObject == nil {
		gateControlObject = new(gateControl)
		gateControlObject.init()
		return gateControlObject
	}

	return gateControlObject
}

func (this * gateControl)init(){
	this.playControlMap = collect.NewMap()
	this.sceneSocketMap = collect.NewMap()
}

func (this * gateControl) Start(){

}

func (this *gateControl) AddPlayControl(playControl* PlayControl) {
	this.playControlMap.Put(playControl.GetId() , playControl)
}

func (this *gateControl) RemovePlayControl(playControl *PlayControl) {
	this.playControlMap.Remove(playControl.GetId())
}

func (this *gateControl) GetPlayControl(uId int64)  *PlayControl{
	playControl := this.playControlMap.Get(uId)
	if playControl == nil {
		return nil
	}
	return playControl.(* PlayControl)
}

func (this * gateControl) GetPlayControlSize(){
	this.playControlMap.Size()
}

func (this * gateControl) GetSceneScoket(sceneId int) * socket.SocketClient{

	sceneSocket := this.sceneSocketMap.Get(sceneId)

	if sceneSocket == nil{

		sceneJson := redis.GetSceneBySceneId(sceneId)

		fmt.Println("sceneJson" , sceneJson)

		if sceneJson == nil {			//如果没有这个场景数据的话
			return nil
		}

		address := sceneJson.GetString("ip") + ":" + sceneJson.GetString("port")

		fmt.Println("Create Scene Socket Client : ", address)
		sceneSocket = socket.NewSocketClient(address , socket.NewStringDeEncode() , func(conn *socket.Connection, msgObject interface{}) {

			defer tools.Exception(nil)

			fmt.Println("Revc Scene msgObject = " , msgObject)

			jsonObject := json.ParseJSONObject(msgObject.(string))

			handleId := jsonObject.GetInt64(protocol.MSG_HANDLE_ID)
			jsonObject.Remove(protocol.MSG_HANDLE_ID)

			playContorl := SingleGateControl().GetPlayControl(handleId)

			if playContorl != nil {
				playContorl.Send(jsonObject.ToJSONString())
			}

			tools.Exception(nil)
		})

		this.sceneSocketMap.Put(sceneId , sceneSocket)
	}

	return sceneSocket.(* socket.SocketClient)
}
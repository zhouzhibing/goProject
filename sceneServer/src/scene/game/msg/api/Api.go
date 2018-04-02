package api

import (
	"sync"
	"gameTools/game/tools/collect"
	"gameTools/game/tools/protocol"
	"gameTools/game/tools/json"
	"gameTools/game/tools/net/sockets/socket"
)

type api struct {	}

var apiObject * api
var sceneLock sync.Mutex
var apiMap * collect.Map

func SingleApi() * api{

	if apiObject != nil {
		return apiObject
	}

	sceneLock.Lock()
	defer sceneLock.Unlock()

	if apiObject != nil {
		return apiObject
	}

	if apiObject == nil {
		apiObject = new(api)
		return apiObject
	}

	return apiObject
}


func ( this * api ) InitApi(){

	apiMap = collect.NewMap()

	apiMap.Put(10 , NewUserApi())
}


func (this *api) DoHandle(connection * socket.Connection , o * json.JSONObject){

	protocolNo := o.GetInt("protocolNo")

	mode := protocol.GetMode(protocolNo)

	modeApi := apiMap.Get(mode).(IApi)

	modeApi.Hanndle(connection , protocolNo ,  o)

}